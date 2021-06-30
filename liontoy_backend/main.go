package main

import (
	"github.com/TechMaster/golang/08Fiber/Basic/database"
	"github.com/TechMaster/golang/08Fiber/Basic/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})
	database.Connect()

	app.Static("/", "./public")
	//ngăn truy cập / request từ cổng khác 3000
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	productRouter := app.Group("/product")
	routes.ConfigProductRouter(&productRouter) //http://localhost:3000/product
	routes.Setup(app)

	app.Listen(":3000")

}
