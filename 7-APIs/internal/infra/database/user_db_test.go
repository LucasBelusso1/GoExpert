package database

import (
	"testing"

	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("dev", "dev@dev", "asdf000")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("dev", "dev@dev", "asdf000")
	userDB := NewUser(db)

	userDB.Create(user)

	userFound, err := userDB.FindByEmail("dev@dev")
	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotEmpty(t, userFound.Password)
}
