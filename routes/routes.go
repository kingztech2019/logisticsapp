package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/9jarider/controllers"
	"github.com/kingztech2019/9jarider/middlewares"
)

func Setup(app *fiber.App) {
	
app.Post("/api/register", controllers.Register)
app.Post("/api/login", controllers.Login)

app.Get("/api/get-blog", controllers.AllPost)
app.Get("/api/activate", controllers.ActivateUser)
app.Post("/api/password-reset-code", controllers.PasswordCodeConfirm)
app.Post("/api/reset-password", controllers.ForgetPassword)

app.Use(middlewares.IsAuthenticated)
app.Get("/api/user", controllers.User)
app.Post("/api/logout", controllers.Logout)
app.Post("/api/users", controllers.CreateUser)
//Users api
app.Get("/api/users", controllers.AllUsers)
app.Get("/api/users/:id", controllers.GetUser)
app.Put("/api/users/:id", controllers.UpdateUser)
app.Delete("/api/users/:id", controllers.DeleteUser)
app.Post("/api/create-blog", controllers.CreateBlog)


//Roles api
app.Get("/api/roles", controllers.AllRoles)
app.Post("/api/roles", controllers.CreateRole)
 app.Get("/api/roles/:id", controllers.GetRole)
app.Put("/api/roles/:id", controllers.UpdateRole)
app.Delete("/api/roles/:id", controllers.DeleteRole)
app.Get("/api/permissions", controllers.AllPermission)

//Adverts API
app.Post("/api/advert", controllers.CreateAdvert)
app.Get("/api/adverts", controllers.AllAdverts)
app.Get("/api/user-advert", controllers.UserAdvert)
app.Put("/api/update-user-advert/:id", controllers.UpdateUserAdvert)
app.Get("/api/advert/:id", controllers.GetAdvert)

app.Post("/api/profile", controllers.CreateProfile)
app.Get("/api/profile", controllers.AllProfile)

//This API is to save image to db
app.Post("/api/upload-image", controllers.Upload)
app.Delete("/api/delete-image", controllers.HandleDeleteImage)
app.Static("/api/uploads", "./uploads")
}