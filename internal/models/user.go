package models

import (
	"errors"
	"time"

	"online_store/internal/authentication"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          uint      `gorm:"primary_key;" json:"id"`
	Fullname    string    `gorm:"size:255;not null;" json:"fullname"`
	Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Password    string    `gorm:"size:255;not null;" json:"password"`
	Email       string    `gorm:"size:255;not null;unique" json:"email"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	UpdatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_date"`
}

func (u *User) Create() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func GetUserByID(id uint) (User, error) {
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}
	user.HiddenPassword()

	return user, nil
}

func (user *User) HiddenPassword() {
	user.Password = ""
}

func VerifyPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := authentication.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}
	return token, nil
}
