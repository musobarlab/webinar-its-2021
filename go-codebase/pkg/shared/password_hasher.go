package shared

import (
	"crypto/sha1"

	pbkdf2 "github.com/wuriyanto48/go-pbkdf2"
)

var (
	// Pbkdf2Hasher password hasher
	Pbkdf2Hasher = pbkdf2.NewPassword(sha1.New, 8, 32, 1000)
)
