package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/aicomylleville/formula-mars_go-api/utils/token"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string  `gorm:"size:255;not null;unique" json:"username"`
	Password string  `gorm:"size:255;not null;" json:"password"`
	Wallet   float64 `gorm:"type:decimal(10,2)" json:"wallet"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return token, nil

}

func (user *User) UpdateWallet() (*User, error) {

	if err := DB.Model(&user).Where(user.ID).Updates(user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) SaveUser() (*User, error) {
	var err error

	err = DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
