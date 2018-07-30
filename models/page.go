package models

import "github.com/jinzhu/gorm"

// Page implementation of what is in the database
type Page struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string
}
