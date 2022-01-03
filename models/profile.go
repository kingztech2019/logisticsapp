package models

import "gorm.io/gorm"

type Profile struct{
	gorm.Model
	 
	Name string `json:"name"`
	Age  string `json:"age"`
	UserID uint  `json:"userid"`
	User User `json:"user";gorm:"foreignkey:UserID"`
	 
	 	 
}