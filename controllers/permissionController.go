package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
)

func AllPermission(c *fiber.Ctx) error {
	var permission []models.Role
	database.DB.Find(&permission)
	return c.JSON(permission)
	
}