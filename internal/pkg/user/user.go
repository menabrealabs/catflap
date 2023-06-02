package user

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/menabrealabs/catflap/internal/pkg/port"
)

type User struct {
	Name  string
	Pass  string
	Ports port.Set
}

// Initialize a reference to a new User
func New(name, password string) (*User, error) {
	pass, err := EncryptPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:  name,
		Pass:  pass,
		Ports: make(port.Set),
	}, nil
}

// Implements the String interface
func (user *User) String() string {
	return user.Name
}

// Update the user name and password
func (user *User) Update(name, raw_pass string) {
	if name != "" {
		user.Name = name
	}

	if raw_pass != "" {
		user.Pass, _ = EncryptPassword(raw_pass)
	}
}

func (user *User) Authenticate(password string) bool {
	return user.Pass == password
}

func EncryptPassword(raw_pass string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(raw_pass))
	if err != nil {
		return "", errors.New("failed to write to Sha256 hash")
	}

	hashed := h.Sum(nil)

	return fmt.Sprintf("%x", hashed), nil
}
