package crypt

import (
	"encoding/base64"
)

// HashPassword transforms a plain string password into a base64 version of a Argon2i hashing
func HashPassword(
	passwd string,
	argon2Bag Argon2Bag,
	salt []byte,
) string {
	return base64.
		StdEncoding.
		EncodeToString(
			Argon2KeyHashBytes(
				[]byte(passwd),
				argon2Bag,
				salt,
			),
		)
}
