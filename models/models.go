package models

import "gorm.io/gorm"

// User struct defines the schema for the table in the database.
type User struct {
	gorm.Model
	// Id int `json:"ID" gorm:"primary_key"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Key string `json:"key"`
}