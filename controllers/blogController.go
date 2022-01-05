package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
)

func CreateBlog(c *fiber.Ctx)  error{
	var blogpost models.Blog
if err:=c.BodyParser(&blogpost);err !=nil{
	return err
}
  
if err:=database.DB.Create(&blogpost).Error;err !=nil{
	 
	c.Status(400)
	 
	return c.JSON(fiber.Map{
	 "message":"Invalid payload",
   })
}
return c.JSON(fiber.Map{
	"message":"Congratulation, Your post is now live",
  })

	
}

func AllPost(c *fiber.Ctx) error {
	page,_:= strconv.Atoi(c.Query("page","1")) 
	 
	return c.JSON(models.Paginate(database.DB,&models.Blog{},page))
	
}