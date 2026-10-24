package runkey

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// actor is what a key stands for in these tests, as an app would have it.
type actor struct {
	UserID string
	KeyID  string
}

// clock is a time the tests move by hand.
type clock struct{ t time.Time }

func (c *clock) now() time.Time       { return c.t }
func (c *clock) pass(d time.Duration) { c.t = c.t.Add(d) }
func newClock() *clock                { return &clock{t: time.Date(2026, 10, 24, 12, 0, 0, 0, time.UTC)} }
func newStore(c *clock) *Store[actor] { return New[actor]("kar_", WithClock(c.now)) }

func Test_A_Key_Stands_For_Its_Value_Until_It_Expires(t *testing.T) {
	c := newClock()
	s := newStore(c)
	secret, _, err := s.Issue(actor{UserID: "u1", KeyID: "k1"}, time.Minute)
	require.NoError(t, err)
	assert.True(t, strings.HasPrefix(secret, "kar_"))
	assert.True(t, s.IsKey(secret))

	got, err := s.Authenticate(secret)
	require.NoError(t, err)
	assert.Equal(t, actor{UserID: "u1", KeyID: "k1"}, got)

	c.pass(time.Minute)
	_, err = s.Authenticate(secret)
	assert.ErrorIs(t, err, ErrInvalidKey)
}

func Test_A_Revoked_Key_Is_Invalid(t *testing.T) {
	s := newStore(newClock())
	secret, revoke, err := s.Issue(actor{UserID: "u1"}, time.Minute)
	require.NoError(t, err)
	revoke()
	revoke() // twice is fine
	_, err = s.Authenticate(secret)
	assert.ErrorIs(t, err, ErrInvalidKey)
}

func Test_An_Unknown_Key_Is_Invalid(t *testing.T) {
	s := newStore(newClock())
	got, err := s.Authenticate("kar_unknown")
	assert.ErrorIs(t, err, ErrInvalidKey)
	assert.Equal(t, actor{}, got)
	assert.False(t, s.IsKey("gak_other"))
}

func Test_RevokeWhere_Ends_The_Matching_Keys_Only(t *testing.T) {
	s := newStore(newClock())
	discord, _, _ := s.Issue(actor{UserID: "u1", KeyID: "discord"}, time.Minute)
	discordAgain, _, _ := s.Issue(actor{UserID: "u2", KeyID: "discord"}, time.Minute)
	slack, _, _ := s.Issue(actor{UserID: "u1", KeyID: "slack"}, time.Minute)

	s.RevokeWhere(func(a actor) bool { return a.KeyID == "discord" })

	for _, secret := range []string{discord, discordAgain} {
		_, err := s.Authenticate(secret)
		assert.ErrorIs(t, err, ErrInvalidKey)
	}
	_, err := s.Authenticate(slack)
	assert.NoError(t, err)
}

func Test_Issuing_Forgets_The_Expired_Keys(t *testing.T) {
	c := newClock()
	s := newStore(c)
	_, _, _ = s.Issue(actor{UserID: "u1"}, time.Minute)
	c.pass(time.Minute)
	_, _, _ = s.Issue(actor{UserID: "u2"}, time.Minute)
	assert.Len(t, s.keys, 1)
}

func Test_Secrets_Are_Distinct(t *testing.T) {
	s := newStore(newClock())
	a, _, _ := s.Issue(actor{UserID: "u1"}, time.Minute)
	b, _, _ := s.Issue(actor{UserID: "u1"}, time.Minute)
	assert.NotEqual(t, a, b)
	// "kar_" and 32 bytes in unpadded base64url
	assert.Len(t, a, len("kar_")+43)
}

func Test_Issue_Fails_Without_Randomness(t *testing.T) {
	s := New[actor]("kar_", WithRandom(bytes.NewReader(nil)))
	_, _, err := s.Issue(actor{UserID: "u1"}, time.Minute)
	assert.Error(t, err)
	assert.False(t, errors.Is(err, ErrInvalidKey))
}
