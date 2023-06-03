package user_test

import (
	"crypto/sha256"
	"testing"

	"github.com/menabrealabs/catflap/internal/pkg/user"
)

const testUserPassphrase = "passphrase"

var testUser = user.User{
	Name: "username",
	Pass: sha256.Sum256([]byte(testUserPassphrase)),
}

func TestUser(t *testing.T) {
	t.Run("should encrypt passphrase when new User reference created", func(t *testing.T) {
		user := user.New(testUser.Name, testUserPassphrase)

		if user.Pass != testUser.Pass {
			t.Errorf("expected: %s, but got: %s", testUser.Pass, user.Pass)
		}
	})

	t.Run("should only print username on Stringer String() implementation", func(t *testing.T) {
		if testUser.String() != testUser.Name {
			t.Errorf("expected: %s, but got: %s", testUser.Name, testUser)
		}
	})
}

func TestUserUpdateName(t *testing.T) {
	t.Run("should update user name", func(t *testing.T) {
		user := user.New(testUser.Name, testUserPassphrase)
		user.UpdateName("luser")

		if user.Name != "luser" {
			t.Errorf("expected: %s, but got: %s", "luser", user.Name)
		}
	})

	t.Run("should return error when name is too short", func(t *testing.T) {
		user := user.New(testUser.Name, testUserPassphrase)
		err := user.UpdateName("no")

		if err == nil {
			t.Error("expected error not thrown on name update")
		}
	})

	t.Run("should not update on error", func(t *testing.T) {
		user := user.New(testUser.Name, testUserPassphrase)
		user.UpdateName("no")

		if user.Name == "no" {
			t.Errorf("expected: %s, but got: %s", testUser.Name, user.Name)
		}
	})
}

func TestUserUpdatePass(t *testing.T) {
	newUserPassphrase := "buttercup-tedious-opinions"
	newUserPass := user.EncryptPassword(newUserPassphrase)

	t.Run("should update user pass", func(t *testing.T) {
		newUser := user.New(testUser.Name, testUserPassphrase)
		newUser.UpdatePass(newUserPassphrase)

		if newUser.Pass != newUserPass {
			t.Errorf("expected: %x, but got: %x", newUserPass, testUser.Pass)
		}
	})

	t.Run("should return error when passphrase is too short", func(t *testing.T) {
		newUser := user.New(testUser.Name, testUserPassphrase)
		err := newUser.UpdatePass("shortpass")

		if err == nil {
			t.Error("expected error not thrown on pass update")
		}
	})

	t.Run("should not update on error", func(t *testing.T) {
		newUser := user.New(testUser.Name, testUserPassphrase)
		newUser.UpdatePass("shortpass")

		if newUser.Pass != testUser.Pass {
			t.Errorf("expected: %x, but got: %x", testUser.Pass, newUser.Pass)
		}
	})
}

func TestUserAuthenticate(t *testing.T) {
	t.Run("should authenticate when given correct passphrase", func(t *testing.T) {
		ok := testUser.Authenticate(testUserPassphrase)

		if !ok {
			t.Error("failed to authenticate correct passphrase")
		}
	})

	t.Run("should not authenticate when given incorrect passphrase", func(t *testing.T) {
		ok := testUser.Authenticate("fiddlepup")

		if ok {
			t.Error("authenticated an incorrect passphrase")
		}
	})
}
