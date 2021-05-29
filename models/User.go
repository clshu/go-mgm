package models

import (
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
)

// User contains user info
type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lastName" bson:"lastName"`
	Password         string `json:"password" bson:"password"`
}

// func NewBook(u User) *User {
// 	return &User{Email: u.Email, FirstName: u.FirstName, LastName: u.LastName, Password: u.Password}}
// }

// Creating : an preop
func (u *User) Creating() error {
	// Call the DefaultModel Creating hook
	if err := u.DefaultModel.Creating(); err != nil {
		return err
	}

	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}
	// We can validate the fields of a model and return an error to prevent a document's insertion.

	return nil
}

// ComparePassword compare hashed and plain text password
func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
