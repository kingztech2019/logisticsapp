package controllers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"gorm.io/gorm"
)



func ActivateUser(c *fiber.Ctx)  error{ 
verify:= c.Query("verify")
var activate models.Activate
 
dataCheck:=database.DB.Where("token = ?", verify).First(&activate)
if errors.Is(dataCheck.Error, gorm.ErrRecordNotFound){
    c.Status(404)
  return c.JSON(fiber.Map{
   "message": "Invalid verification link",
   })
}
//If token is expired triggered this
timeCreated:=activate.CreatedAt
expiredTime:=timeCreated.Add(2 * time.Hour)
compareDate:=time.Now().After(expiredTime) 
if compareDate {
	c.Status(404)
	return c.JSON(fiber.Map{
	 "message": "Token is already expired",
	 })
	
}else{
	//if user has already been activated triggered this
	database.DB.Where("token = ? AND used=?" , verify, 1).First(&activate)
	if activate.Used {
	 return c.JSON(fiber.Map{
		"message": "Your account has already been activated. Proceed to login",
		})
	
	}
	dataCheck.Update("used",1)
	
}

 
return c.JSON(fiber.Map{
	"userId": activate.UserID,
	})
   
}
 
	