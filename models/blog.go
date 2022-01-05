package models

import "gorm.io/gorm"

type Blog struct{
	 
	gorm.Model
	Title string `json:"title"`
	Desc  string `json:"desc"`
	UserID uint  `json:"userid"`
	User User `json:"user";gorm:"foreignkey:UserID"`
	Image string `json:"image"`
	 
	 	 
}

func (blog *Blog)  Count(db *gorm.DB) int64{
	var total int64
	db.Model(&Blog{}).Count(&total)
	return total
		
	}
	
	func (blog *Blog)Take(db *gorm.DB,limit int, offset int) interface{}  {
		var blogs []Blog
		db.Offset(offset).Limit(limit).Preload("User").Find(&blogs)
		return blogs
		
	}