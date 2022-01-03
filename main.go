package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/kingztech2019/9jarider/database"
	//"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/routes"
)

func main() {
	 
	database.Connect()
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	port:=os.Getenv("PORT")
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
 routes.Setup(app)
//   MailChan:= make(chan models.MailData)
//    defer close(MailChan)

 app.Listen(port)
}