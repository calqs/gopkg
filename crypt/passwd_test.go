package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I_Can_Hash_A_Password(t *testing.T) {
	defaultBag := GetBasicArgon2Conf()
	// to verify this, with cli tools:
	// $> argon2 12345678 -t 3 -k 32768 -p 4 -l 32
	// 	cabane
	// Type:           Argon2i
	// Iterations:     3
	// Memory:         32768 KiB
	// Parallelism:    4
	// Hash:           c00a0bf1aafdd3819d09c93dd1a674ebe16727665d5e3adde45a59319f87ca27
	// Encoded:        $argon2i$v=19$m=32768,t=3,p=4$MTIzNDU2Nzg$wAoL8ar904GdCck90aZ06+FnJ2ZdXjrd5FpZMZ+Hyic
	// 0.136 seconds
	// Verification ok
	// $> echo c00a0bf1aafdd3819d09c93dd1a674ebe16727665d5e3adde45a59319f87ca27 \
	// | xxd -r -p \
	// | base64
	//
	// wAoL8ar904GdCck90aZ06+FnJ2ZdXjrd5FpZMZ+Hyic=
	assert.Equal(t, "wAoL8ar904GdCck90aZ06+FnJ2ZdXjrd5FpZMZ+Hyic=", HashPassword("cabane\n", defaultBag, []byte("12345678")))
}
