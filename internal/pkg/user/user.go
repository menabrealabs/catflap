package user

import (
	"crypto/sha256"
	"errors"

	"github.com/menabrealabs/catflap/internal/pkg/port"
)

// The Sha256 checksum of the user passphrase in a 32 byte array
type Checksum [sha256.Size]byte

// A User record stores username, a Sha256 checksum of their
// passphrase, and a set of ports the user can access.
type User struct {
	Name  string
	Pass  Checksum
	Ports port.Set
}

// Initialize a reference to a new User.
func New(name, raw_passphrase string) (*User, error) {
	pass, err := EncryptPassword(raw_passphrase)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:  name,
		Pass:  pass,
		Ports: make(port.Set),
	}, nil
}

// Implements the String interface.
func (user *User) String() string {
	return user.Name
}

// Update the user name and password.
func (user *User) Update(name, raw_passphrase string) {
	if name != "" {
		user.Name = name
	}

	if raw_passphrase != "" {
		user.Pass, _ = EncryptPassword(raw_passphrase)
	}
}

// Authenticates plaintext passphrase against the stored encrypted pass.
func (user *User) Authenticate(raw_passphrase string) (bool, error) {
	pass, err := EncryptPassword(raw_passphrase)
	if err != nil {
		return false, err
	}
	return (user.Pass == pass), nil
}

// Encrypt a raw plaintext passphrase into a Sha256 hashed pass.
func EncryptPassword(raw_passphrase string) (Checksum, error) {
	h := sha256.New()
	_, err := h.Write([]byte(raw_passphrase))
	if err != nil {
		return Checksum{}, errors.New("failed to write to Sha256 hash")
	}

	hashed := Checksum(h.Sum(nil))

	return hashed, nil
}
