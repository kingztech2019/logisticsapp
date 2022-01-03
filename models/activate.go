package models

import (
	"gorm.io/gorm"
)

type Activate struct{
	gorm.Model
	Token string `json:"token"`
	Expired bool
	Used bool
	UserID uint  `json:"userid"`
 
	 
	 	 
}