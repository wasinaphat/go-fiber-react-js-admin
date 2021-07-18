package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wasinaphatlilawatthananan/go-admin/controllers"
	"github.com/wasinaphatlilawatthananan/go-admin/middleware"
)

func Setup(app *fiber.App) {

	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)
	app.Put("api/users/info", controllers.UpdateInfo)
	app.Put("api/users/password", controllers.UpdatePassword)
	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Get("api/users", controllers.AllUsers)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)
	app.Post("api/users", controllers.CreateUser)

	app.Get("api/roles", controllers.AllRoles)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)
	app.Post("api/roles", controllers.CreateRole)

	app.Get("api/permissions", controllers.AllPermissions)

	app.Get("api/products", controllers.AllProducts)
	app.Get("api/products/:id", controllers.GetProduct)
	app.Put("api/products/:id", controllers.UpdateProduct)
	app.Delete("api/products/:id", controllers.DeleteProduct)
	app.Post("api/products", controllers.CreateProduct)

	app.Post("api/upload", controllers.Upload)
	app.Static("/api/uploads","./uploads")

	app.Get("api/orders", controllers.AllOrders)
	app.Post("api/export", controllers.Export)
	app.Get("api/chart", controllers.Chart)
}
