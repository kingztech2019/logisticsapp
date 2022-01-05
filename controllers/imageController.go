package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error  {
	form, err:=c.MultipartForm()
	if err != nil {
		return err
	}
	files:=form.File["image"]
	filename:=""

	for _,file:=range files{
		filename=file.Filename
 
		if err:=c.SaveFile(file, "./uploads/"+filename);  err!=nil{
			return nil
		}	
	
}
 return c.JSON(fiber.Map{
	"url":"https://test-logistics-app.herokuapp.com/api/uploads/"+filename,
	 
})
}
func HandleDeleteImage(c *fiber.Ctx) error {
    // extract image name from params
    imageName := c.Query("imageName")
	fmt.Println(imageName)

    // delete image from ./images
    err := os.Remove("./uploads/"+imageName)
    if err != nil {
        log.Println(err)
		c.Status(500)
        return c.JSON(fiber.Map{"status": 500, "message": "Unable to delete image",})
    }

    return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", })
}