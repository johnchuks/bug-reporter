package models

import (
	"log"
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// User definition
type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100)" validate:"required"  json:"firstName"`
	LastName string `gorm:"type:varchar(100)"   json:"lastName"`
	UserName string `gorm:"type:varchar(100);not null;unique" validate:"required unique" json:"userName"`
	Password string `gorm:"type:varchar(100);not null" validate:"required" json:"password"`
}

// Create a new user object
func (u *User) Create(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Update an existing user
func (u *User) Update(id int,  db *gorm.DB) (*User, error) {
	var err error
	user := &User{
		FirstName: u.FirstName,
		LastName: u.LastName,
		UserName: u.UserName,
	}
	err = db.Debug().Table("users").Where("id = ?", id).Updates(user).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// GetByUserName gets a user object with their username
func (u *User) GetByUserName(username string, db *gorm.DB) (*User, error) {
	var err error
	user := &User{}
	err = db.Debug().Table("users").Where("username = ?", username).First(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

//Get a user by their id
func (u *User) Get(id int, db *gorm.DB) (*User, error) {
	var err error
	user := &User{}
	err = db.Debug().Table("users").Where("id = ?", id).First(user).Error

	if err != nil {
		return &User{}, err
	}
	return user, nil
}

var validate *validator.Validate

// BeforeCreate hook for checking if data is valid before saving it
func (u *User) BeforeCreate() (err error) {
	if !u.IsValid() {
	  err = errors.New("can't save invalid data")
	  return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return err
	}
	u.Password = string(hash)
	return
}

// IsValid checks if the user input is valid before saving it
func (u *User) IsValid() bool {
	err := validate.Struct(u)
	if err != nil {
		return false
	}
	return true
}