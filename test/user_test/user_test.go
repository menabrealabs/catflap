package user_test

import (
	"crypto/sha256"
	"testing"

	"github.com/menabrealabs/catflap/internal/pkg/user"
)

func TestUser(t *testing.T) {
	t.Run("should encrypt passphrase when new User reference created", func(t *testing.T) {
		user := user.New("username", "passphrase")

		expected := sha256.Sum256([]byte("passphrase"))

		if user.Pass != expected {
			t.Errorf("expected: %s, but got: %s", expected, user.Pass)
		}
	})

}
