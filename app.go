package main

import (
	"github.com/TechMaster/golang/08Fiber/Repository/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	userRouter := app.Group("/user")
	routes.ConfigUserRouter(&userRouter) //http://localhost:3000/user
	routes.Setup(app)
	productRouter := app.Group("/product")
	routes.ConfigProductRouter(&productRouter) //http://localhost:3000/product
	routes.Setup(app)
	categoryRoute := app.Group("/category")
	routes.ConfigCategoryRouter(&categoryRoute) //http://localhost:3000/product
	routes.Setup(app)
	app.Listen(":3000")
}
