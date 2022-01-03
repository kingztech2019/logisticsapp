package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
)


func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Find(&roles)
	return c.JSON(roles)
	
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role
if err:=c.BodyParser(&role);err !=nil{
	return err
} 
database.DB.Create(&role)
return c.JSON(role)
 
	
}

func GetRole(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id")) 
	role:=models.Role{
		Id:uint(id),
	}
	database.DB.Find(&role)
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	role:=models.Role{
		Id:uint(id),
	} 
	if err:=c.BodyParser(&role);err !=nil{
		return err
	}

	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
	
}

func DeleteRole(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	role:=models.Role{
		Id:uint(id),
	} 
	database.DB.Delete(&role)
	return c.JSON(fiber.Map{
		"message":"Account deleted successfully",
	  })
	  
	
}