package user

import (
	"github.com/menabrealabs/catflap/internal/pkg/port"
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

// Update the user name and password.
func (user *User) Update(name, raw_passphrase string) {
	if name != "" {
		user.Name = name
	}

	if raw_passphrase != "" {
		user.Pass = EncryptPassword(raw_passphrase)
	}
}

// Authenticates plaintext passphrase against the stored encrypted pass.
func (user User) Authenticate(raw_passphrase string) bool {
	pass := EncryptPassword(raw_passphrase)
	return (user.Pass == pass)
}
