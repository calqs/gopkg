// Package runkey keeps short-lived bearer keys in memory. Each key stands for
// a value (a user, a principal…) while a job runs, e.g. an AI agent calling
// an API back as the user, and ends with it.
package runkey

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// secretBytes is the randomness of a secret: 256 bits.
const secretBytes = 32

// ErrInvalidKey: the key is unknown, revoked or expired.
var ErrInvalidKey = errors.New("runkey: invalid key")

// Store keeps the live keys of one process, by their secret's hash: a secret
// is never kept, and a restart ends every key.
type Store[T any] struct {
	mu     sync.Mutex
	prefix string
	keys   map[[sha256.Size]byte]entry[T]
	now    func() time.Time
	random io.Reader
}

type entry[T any] struct {
	value     T
	expiresAt time.Time
}

type config struct {
	now    func() time.Time
	random io.Reader
}

// Option changes a Store's clock or randomness, for tests.
type Option func(*config)

// WithClock makes the store tell the time with now (time.Now by default).
func WithClock(now func() time.Time) Option {
	return func(c *config) { c.now = now }
}

// WithRandom makes the store draw its secrets from random (crypto/rand's
// Reader by default).
func WithRandom(random io.Reader) Option {
	return func(c *config) { c.random = random }
}

// New is a store of keys starting with prefix ("kar_"), which tells them from
// the other secrets an API takes.
func New[T any](prefix string, opts ...Option) *Store[T] {
	c := config{now: time.Now, random: rand.Reader}
	for _, opt := range opts {
		opt(&c)
	}
	return &Store[T]{prefix: prefix, keys: map[[sha256.Size]byte]entry[T]{}, now: c.now, random: c.random}
}

// IsKey: secret has the store's prefix. It may still be invalid.
func (s *Store[T]) IsKey(secret string) bool {
	return strings.HasPrefix(secret, s.prefix)
}

// Issue makes a secret standing for value until ttl passes or revoke is
// called. Calling revoke more than once is fine.
func (s *Store[T]) Issue(value T, ttl time.Duration) (secret string, revoke func(), err error) {
	secret, err = s.newSecret()
	if err != nil {
		return "", nil, err
	}
	hash := sha256.Sum256([]byte(secret))
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sweep()
	s.keys[hash] = entry[T]{value: value, expiresAt: s.now().Add(ttl)}
	return secret, func() { s.revoke(hash) }, nil
}

// Authenticate is the value a live secret stands for; ErrInvalidKey when it
// is unknown, revoked or expired.
func (s *Store[T]) Authenticate(secret string) (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.keys[sha256.Sum256([]byte(secret))]
	if !ok || !s.now().Before(e.expiresAt) {
		var zero T
		return zero, ErrInvalidKey
	}
	return e.value, nil
}

// RevokeWhere ends every live secret whose value matches, e.g. all the keys
// of a user who unlinked an integration.
func (s *Store[T]) RevokeWhere(match func(T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for hash, e := range s.keys {
		if match(e.value) {
			delete(s.keys, hash)
		}
	}
}

func (s *Store[T]) newSecret() (string, error) {
	raw := make([]byte, secretBytes)
	if _, err := io.ReadFull(s.random, raw); err != nil {
		return "", fmt.Errorf("runkey: drawing a secret: %w", err)
	}
	return s.prefix + base64.RawURLEncoding.EncodeToString(raw), nil
}

func (s *Store[T]) revoke(hash [sha256.Size]byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.keys, hash)
}

// sweep forgets the expired keys, so a job that never revoked its key does
// not leave it behind. Called with the lock held.
func (s *Store[T]) sweep() {
	now := s.now()
	for hash, e := range s.keys {
		if !now.Before(e.expiresAt) {
			delete(s.keys, hash)
		}
	}
}
