package models

import (
	"strings"

	"github.com/clshu/go-mgm/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User contains user info
type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Email            string             `json:"email" bson:"email"`
	FirstName        string             `json:"firstName" bson:"firstName"`
	LastName         string             `json:"lastName" bson:"lastName"`
	Password         string             `json:"password" bson:"password"`
	TempPassword     utils.TempPassword `json:"tempPassword" bson:"tempPassword"`
}

// UserViewer is a return type with partial user info
type UserViewer struct {
	ID        primitive.ObjectID `json:"id"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
}

// UserAuth is an auth payload with partial user info and token
type UserAuth struct {
	User  UserViewer `json:"user"`
	Token string     `json:"token"`
}

// Creating : an preop to create hashed password and lower case email
func (u *User) Creating() error {
	// Call the DefaultModel Creating hook
	if err := u.DefaultModel.Creating(); err != nil {
		return err
	}

	if u.Email != "" {
		u.Email = strings.ToLower(u.Email)
	}
	if u.Password != "" {
		hash, err := utils.CreateHashedPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}

	return nil
}
