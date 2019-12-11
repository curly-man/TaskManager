package models

import (
	// "crypto/sha1"
	// "bytes"
)

type User struct {
	Username string
	Login    string
	Password []byte
}