package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/util"
)

func AllAdverts(c *fiber.Ctx) error {
	var adverts []models.Advert
	database.DB.Preload("User").Find(&adverts)
	return c.JSON(fiber.Map{
		"adverts":adverts,
	  })
	
}

func CreateAdvert(c *fiber.Ctx) error {
	var advert models.Advert
if err:=c.BodyParser(&advert);err !=nil{
	return err
}
  
if err:=database.DB.Create(&advert).Error;err !=nil{
	 
	c.Status(400)
	 
	return c.JSON(fiber.Map{
	 "message":"Invalid payload",
   })
}
return c.JSON(advert)
 
	
}




func GetAdvert(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id")) 
	
	advert:=models.Advert{} 
	   database.DB.Preload("User").First(&advert, id)
	 
	 
	// if errors.Is(data.Error, gorm.ErrRecordNotFound){
	// 		c.Status(404)
	// 	return c.JSON(fiber.Map{
	// 	 "message": "User not found",
	//    })
	// }
	// if data.RowsAffected==0 {
	// 	c.Status(404)
	// 	return c.JSON(fiber.Map{
	// 	 "message": "User not found",
	//    })
	// }
	 
	 
	return c.JSON(fiber.Map{
		"singleAdvert":advert,
	  })
	 
	 
}

//This function is to get the authenticated user
func UserAdvert(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
    id, _:= util.ParseJwt(cookie)
	//id,_:=strconv.Atoi(c.Params("id")) 
	
	var advert []models.Advert
	database.DB.Model(&advert).Where("user_id=?", id).Preload("User").Find(&advert)
	//  database.DB.Preload("User").Find(&profile)
	if len(advert)<=0 {
		c.Status(404)
		return c.JSON(fiber.Map{
		 "message": "No advert available for this user",
	   })
		
	}
	
	return c.JSON(fiber.Map{
		"userAdvert":advert,
	  })
	
}

//This function is for user to be able to edit their advert
func UpdateUserAdvert(c *fiber.Ctx) error {
	id,_:=strconv.Atoi(c.Params("id"))
	advert:=models.Advert{
		Id:uint(id),
	} 
	if err:=c.BodyParser(&advert);err !=nil{
		return err
	}

	if err:= database.DB.Model(&advert).Updates(advert).Error; err!=nil{
		return c.JSON(fiber.Map{
			"message":"Error in updating the product",
		  })
	}
	return c.JSON(fiber.Map{
		"message":"Advert updated successfully",
	  })
	
}
