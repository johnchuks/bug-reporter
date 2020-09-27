package models

import (
	"strings"
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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return err
	}
	u.Password = string(hash)
	return
}

// CheckInput checks if the user input is valid before saving it
func (u *User) CheckInput() (err error) {
	err = validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) checkPassword(password string) (isPasswordValid bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// Strip removes all whitespaces from the request body
func (u *User) Strip() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.UserName = strings.TrimSpace(u.UserName)
	u.Password = strings.TrimSpace(u.Password)
}