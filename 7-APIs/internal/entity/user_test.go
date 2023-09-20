package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Dev", "dev@dev.com", "asdf000")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Dev", user.Name)
	assert.Equal(t, "dev@dev.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Dev", "dev@dev.com", "asdf000")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("asdf000"))
	assert.False(t, user.ValidatePassword("asdf00"))
	assert.NotEqual(t, "asdf000", user.Password)
}
