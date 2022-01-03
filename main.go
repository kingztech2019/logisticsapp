package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/kingztech2019/9jarider/database"
	//"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/routes"
)

func main() {
	 
	database.Connect()
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
 routes.Setup(app)
//   MailChan:= make(chan models.MailData)
//    defer close(MailChan)

 app.Listen(":3000")
}