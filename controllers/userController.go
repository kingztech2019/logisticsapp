package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"gorm.io/gorm"
)


func AllUsers(c *fiber.Ctx) error {
	page,_:= strconv.Atoi(c.Query("page","1")) 
	 
	return c.JSON(models.Paginate(database.DB,&models.User{},page))
	
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
if err:=c.BodyParser(&user);err !=nil{
	return err
}
 
user.SetPassword("123456")
database.DB.Create(&user)
return c.JSON(user)
 
	
}

func GetUser(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id")) 
	user:=models.User{
		Id:uint(id),
	}

var activate models.Activate
	data:=database.DB.Preload("Role").Preload("Activate").First(&user, id)
	 
	if errors.Is(data.Error, gorm.ErrRecordNotFound){
			c.Status(404)
		return c.JSON(fiber.Map{
		 "message": "User not found",
	   })
	}
	// if data.RowsAffected==0 {
	// 	c.Status(404)
	// 	return c.JSON(fiber.Map{
	// 	 "message": "User not found",
	//    })
	// }
	  
	database.DB.Where("used=? and user_id=?" , 1, id).Find(&activate)
	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"user":user, 
			"activate":activate.Used,
		},
	  })
	  return c.JSON(user)
	 
	 
}

func UpdateUser(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	user:=models.User{
		Id:uint(id),
	} 
	if err:=c.BodyParser(&user);err !=nil{
		return err
	}

	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
	
}

func DeleteUser(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	user:=models.User{
		Id:uint(id),
	} 
	database.DB.Delete(&user)
	return c.JSON(fiber.Map{
		"message":"Account deleted successfully",
	  })
	  
	
}