package models

import (
	"errors"
	"fmt"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);" json:"email"`
	Password  string `gorm:"size:100; not null" json:"password"`
	FirstName string `gorm:"size:100;not null" json:"firstname"`
	LastName  string `gorm:"size:100;not null" json:"lastname"`
}


func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(passwordBytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}


func (u * User) Prepare() {
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
}

func (u *User) BeforeSave() error{
	hashedPassword, err := HashPassword(u.Password)
	if err != nil{
		return err
	}

	u.Password = hashedPassword

	return nil
}


func (u *User) Validate(action string) error {
	switch action {
	case "login":
		if u.Email == "" {
			return errors.New("email is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
	default:
		if u.Email == "" {
			return errors.New("email is required")
		}
		if u.Password == "" {
			fmt.Println("password is null")
			return errors.New("password is required")
		}
		if u.FirstName == "" {
			return errors.New("first name is required")
		}
		if u.LastName == "" {
			return errors.New("last name is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email format is incorrect")
		}
	}

	return nil
}


func (u *User) Save(db *gorm.DB) (*User, error) {
	result := db.Create(&u)
	if result.Error != nil {
		return &User{}, result.Error
	}

	return u, nil

}


func (u *User) GetUser(db *gorm.DB) (*User, error) {
	user := &User{}
	if err := db.Debug().Table("users").Where("email = ?", u.Email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}




