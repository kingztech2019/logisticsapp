package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/util"
	"gorm.io/gorm"
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

func DetailBlog(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id")) 
	// blog:=models.Blog{
	// 	Id:uint(id),
		
		 
	// }
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data":blogpost,
	  })

	
}

func UpdateBlog(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	blog:=models.Blog{
		Id:uint(id),
	} 
	if err:=c.BodyParser(&blog);err !=nil{
		return err
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(blog)
	
}

func UniqueBlog(c *fiber.Ctx) error  {
	cookie := c.Cookies("jwt")
    id, _:= util.ParseJwt(cookie)
	var blog []models.Blog
	fmt.Println(id)
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	 
	return c.JSON(blog)
	
	
}

func DeleteBlog(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	blog:=models.Blog{
		Id:uint(id),
	} 
	deleteQuery:=database.DB.Delete(&blog)
	  if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound){
    c.Status(400)
  return c.JSON(fiber.Map{
   "message": "Opps! Post not found",
   })
}
	return c.JSON(fiber.Map{
		"message":"Post deleted successfully",
	  })
	  
	
}