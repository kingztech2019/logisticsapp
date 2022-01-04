package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/database"
	"github.com/kingztech2019/9jarider/models"
	"github.com/kingztech2019/9jarider/render"
	"github.com/kingztech2019/9jarider/util"
	"gorm.io/gorm"
)

//This function validate the email address
func validateEmail(email string) bool {
  Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
  return Re.MatchString(email)
}

//This function is to generate random token for user
func generateToken() (string, error) {
  b := make([]byte, 35)
  _, err := rand.Read(b)
  if err != nil {
      return "", err
  }
  return base64.URLEncoding.EncodeToString(b), nil
}

//This function is to generate passord reset token for users
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//This function is to register the user
func Register(c *fiber.Ctx) error{
  var data map[string]interface{}
  var userdata models.User

 if err:=c.BodyParser(&data); err != nil {
    fmt.Println("Unable to parse body")
  }
  //Check if the password length is more than 6
  if len(data["password"].(string))  <= 6 {
    c.Status(400)
    return c.JSON(fiber.Map{
    "message": "Password must be more than 6 character",
  })
     
  }

  // Check if email address already exist in database
  database.DB.Where("email=?",  strings.TrimSpace(data["email"].(string))).First(&userdata)
  if userdata.Id !=0 {
    c.Status(400)
   return c.JSON(fiber.Map{
    "message": "Email already exist",
  })
  
  }
   
//check if the email address is valid email
   if !validateEmail(strings.TrimSpace(data["email"].(string))) {
    c.Status(400)
   return c.JSON(fiber.Map{
    "message": "Invalid email address",
  })
  
     
   }
   //Hash the password
 
  user:= models.User{
    
    FirstName:  data["first_name"].(string),
    LastName:   data["last_name"].(string),
    //Phone:int(data["phone"].(float64)),
    Phone:  data["phone"].(string),
    // RoleId: 1, 
    Email:  strings.TrimSpace(data["email"].(string)),



  }
  
  //This function triggered the verification email
 generatedToken,_:= generateToken()
 
  //render.SendEmail(data["first_name"].(string),data["email"].(string),generatedToken)
  
  user.SetPassword(data["password"].(string))
  err:=database.DB.Create(&user)
  if err != nil {
    log.Println(err)
  }

  activate:=models.Activate{
    Token:generatedToken ,
    Used: false,
    UserID: user.Id,
    Expired: false,
    
  }
  database.DB.Create(&activate)
  return c.JSON(fiber.Map{
    "user": user,
    "message":"You have successfully register",
     

  })
   
     
  }

  //-------------This is login function----------------
  func Login(c *fiber.Ctx) error {
    var data map[string]string 
  if err:=c.BodyParser(&data); err != nil {
    fmt.Println("Unable to parse body")
  }
  var user models.User
  database.DB.Where("email=?", data["email"]).First(&user)

  
  if user.Id ==0{
    c.Status(404)
     return c.JSON(fiber.Map{
      "message": "Invalid Email Address",
    })
    
    
  }
  var activate models.Activate
  dataCheck:=database.DB.Where("user_id = ? AND used = ?",user.Id , 1).First(&activate)
 
  if errors.Is(dataCheck.Error, gorm.ErrRecordNotFound){
    c.Status(400)
  return c.JSON(fiber.Map{
   "message": "Please kindly  activate your account",
   })
}

 if err:= user.ComparePassword(data["password"]); err!=nil{
   c.Status(400)
  return c.JSON(fiber.Map{
     "message":"incorrect password",
   })
   
 }

 // Issuer: strconv.Itoa(int(user.Id)),
  
 token,err:=util.GenerateJwt(strconv.Itoa(int(user.Id)),) 
 if err != nil {
     c.Status(fiber.StatusInternalServerError)
    return nil
 }
 cookie := fiber.Cookie{
   Name:"jwt",
   Value:token,
   Expires:time.Now().Add(time.Hour*24), // 1 day
   HTTPOnly: true,
  //  SameSite: "None",
  //  Secure:   false,

    
   
 }
 c.Cookie(&cookie)
   return c.JSON(fiber.Map{
     "messsage":"success",
     
   })
  
    
  }
  type Cliams struct{
    jwt.StandardClaims
  }
  //This is get auth-user function
  func User(c *fiber.Ctx) error  {
    cookie := c.Cookies("jwt")
    id, _:= util.ParseJwt(cookie)

    
     var user models.User
     database.DB.Where("id=?", id).First(&user)

    return c.JSON(user)
    
  }

  //This function triggered the forget password reset code
  func PasswordCodeConfirm(c *fiber.Ctx) error {
    var data map[string]string
    
    if err:=c.BodyParser(&data); err != nil {
      fmt.Println("Unable to parse body",err)
    }

    if !validateEmail(strings.TrimSpace(data["email"])) {
      c.Status(400)
     return c.JSON(fiber.Map{
      "message": "Invalid email address",
    })
    
       
     }
      
     var user models.User
     database.DB.Where("email=?", strings.TrimSpace(data["email"])).Find(&user)
        
     if user.Id ==0{
       c.Status(404)
        return c.JSON(fiber.Map{
         "message": "User not found",
       })
       
       
     }
     val, _ := randomHex(4)
     var confirmCode=models.PasswordToken{
       Token: val,
       UserID: user.Id,
       Used: false,
     }
    //  timeCreated:=confirmCode.CreatedAt
    // expiredTime:=timeCreated.Add(2 * time.Hour)
    // compareDate:=time.Now().After(expiredTime)
    // if compareDate{
    //   database.DB.Where("user_id = ?", user.Id).Delete(&confirmCode)
    //   c.Status(404)
    //   return c.JSON(fiber.Map{
    //    "message": "Password reset code is expired. ",
    //  })

    // }

     render.SendEmailToken(user.FirstName,user.Email,val)
     database.DB.Create(&confirmCode)

       return c.JSON(&confirmCode)

    
  }

  // This function update the password 
  func ForgetPassword(c *fiber.Ctx) error {
    var data map[string]string
    
    if err:=c.BodyParser(&data); err != nil {
      fmt.Println("Unable to parse body",err)
    }
    var resetpassword models.PasswordToken
    database.DB.Where("token=?", strings.TrimSpace(data["token"])).Find(&resetpassword)
    if resetpassword.UserID ==0{
      c.Status(404)
       return c.JSON(fiber.Map{
        "message": "Invalid Token",
      })
      
      
    }
      timeCreated:=resetpassword.CreatedAt
    expiredTime:=timeCreated.Add(2 * time.Hour)
    compareDate:=time.Now().After(expiredTime)
    if compareDate{
      //database.DB.Where("user_id = ?", user.Id).Delete(&confirmCode)
      c.Status(404)
      return c.JSON(fiber.Map{
       "message": "Password reset code is expired. ",
     })

    }
    database.DB.Model(&resetpassword).Where("token = ?", data["token"]).Update("used", 1)
    cookie:=c.Cookies("jwt")
    id,_ :=util.ParseJwt(cookie)
    user:= models.User{}
    //Check if the password length is more than 6
  if len(data["password"])  <= 6 {
    c.Status(400)
    return c.JSON(fiber.Map{
    "message": "Password must be more than 6 character",
  })
     
  }
    user.SetPassword(data["password"])
    database.DB.Model(&user).Where("id=?", id).Updates(user)
    database.DB.Where("used = ?", 1).Delete(&resetpassword)

    return c.JSON(fiber.Map{
      "message": "Password reset successfully ",
    })


  }


//This is logout function
  func Logout(c *fiber.Ctx) error {
    cookie := fiber.Cookie{
      Name:"jwt",
      Value:"",
      Expires:time.Now().Add(-time.Hour), // 1 day
      HTTPOnly: true,
       
      
    }
    c.Cookie(&cookie)
    return c.JSON(fiber.Map{
      "message":"You logout successfully",
    })
   
  }