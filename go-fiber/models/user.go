package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	Age       int    `json:"age"`
	Education string `json:"education"`
}



