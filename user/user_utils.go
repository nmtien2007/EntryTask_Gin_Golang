package user

import (
	"source_apis/core"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CheckLogin(username string, password string, db *gorm.DB) (string, error) {
	var err error
	user := User{}

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil
	}

	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := core.GenerateToken(int(user.Id))

	if err != nil {
		return "", err
	}

	return token, nil
}
