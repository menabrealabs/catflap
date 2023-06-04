package user

import (
	"errors"

	"github.com/menabrealabs/catflap/internal/port"
)

// A User record stores username, a Sha256 checksum of their
// passphrase, and a set of ports the user can access.
type User struct {
	Name  string
	Pass  Checksum
	Ports port.Set
}

// Initialize a reference to a new User.
func New(name, raw_passphrase string) *User {
	return &User{
		Name:  name,
		Pass:  EncryptPassword(raw_passphrase),
		Ports: make(port.Set),
	}
}

// Implements the String interface.
func (user User) String() string {
	return user.Name
}

func (user *User) UpdateName(name string) error {
	if len(name) < 3 {
		return errors.New("user name must be 3 or more characters")
	}

	user.Name = name
	return nil
}

func (user *User) UpdatePass(raw_passphrase string) error {
	if len(raw_passphrase) < 16 {
		return errors.New("passphrase must be 16 characters or more")
	}

	user.Pass = EncryptPassword(raw_passphrase)
	return nil
}

// Update the user name and password.
func (user *User) Update(name, raw_passphrase string) {
	if name != "" {
		user.Name = name
	}

	if raw_passphrase != "" {
		user.Pass = EncryptPassword(raw_passphrase)
	}
}

// Authenticates plaintext passphrase login against the stored encrypted pass.
func (user User) Authenticate(raw_passphrase string) bool {
	pass := EncryptPassword(raw_passphrase)
	return (user.Pass == pass)
}
