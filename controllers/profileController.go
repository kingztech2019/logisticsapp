package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/util"
)

func CreateProfile(c *fiber.Ctx) error {
	var profile models.Profile
if err:=c.BodyParser(&profile);err !=nil{
	return err
}
  
database.DB.Create(&profile)
return c.JSON(profile)
 
	
}

// func User(c *fiber.Ctx) error  {
//     cookie := c.Cookies("jwt")
//     id, _:= util.ParseJwt(cookie)

    
//      var user models.User
//      database.DB.Where("id=?", id).First(&user)

//     return c.JSON(user)
    
//   }

func AllProfile(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
    id, _:= util.ParseJwt(cookie)
	
	var profile models.Profile
	database.DB.Where("user_id=?", id).Preload("User").First(&profile)
	//  database.DB.Preload("User").Find(&profile)
	return c.JSON(profile)
	
}