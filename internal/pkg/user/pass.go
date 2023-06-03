package user

import (
	"crypto/sha256"
	"encoding/hex"
)

// The Sha256 checksum of the user passphrase in a 32 byte array.
type Checksum [sha256.Size]byte

// Encode the Checksum into a string of hex values.
func (sum Checksum) HexString() string {
	return hex.EncodeToString(sum[:])
}

// Encrypt a raw plaintext passphrase into a Sha256 hashed pass.
func EncryptPassword(raw_passphrase string) Checksum {
	return Checksum(sha256.Sum256([]byte(raw_passphrase)))
}
