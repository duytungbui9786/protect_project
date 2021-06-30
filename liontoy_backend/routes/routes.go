package routes

import (
	"github.com/TechMaster/golang/08Fiber/Basic/controllers"
	"github.com/gofiber/fiber/v2"
)

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controllers.GetAllProduct)

	(*router).Get("/boy", controllers.GetAllProductBoy)

	(*router).Get("/girl", controllers.GetAllProductGirl)

	(*router).Get("/limit", controllers.GetAllProductLimit)

	(*router).Get("/hot", controllers.GetAllProductHot)

	(*router).Get("/feature", controllers.GetAllProductFeature)

	(*router).Get("/new", controllers.GetAllProductNew)

	(*router).Get("/relate", controllers.Get5ProductRelate)

	(*router).Get("/get/:id", controllers.GetProductById)

	(*router).Delete("/:id", controllers.DeleteProductById)

	(*router).Post("", controllers.CreateProduct)

	(*router).Put("", controllers.UpdateProduct)
	//image
}
func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//check user auth login
	// app.Use(middlewares.IsAuthenticated)
	app.Post("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	//user
	app.Get("/api/users", controllers.AllUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)

}
