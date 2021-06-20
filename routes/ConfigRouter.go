package routes

import (
	"github.com/TechMaster/golang/08Fiber/Repository/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigUserRouter(router *fiber.Router) {
	//Return all books
	(*router).Get("/", controller.GetAllUser)

	(*router).Get("/:id", controller.GetUserById)

	(*router).Delete("/:id", controller.DeleteUserById)

	(*router).Post("", controller.CreateUser)

	(*router).Put("", controller.UpdateUser)
}
func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllProduct)

	(*router).Get("/:id", controller.GetProductById)

	(*router).Delete("/:id", controller.DeleteProductById)

	(*router).Post("", controller.CreateProduct)

	(*router).Put("", controller.UpdateProduct)
	//image
}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Get("/", controller.GetAllCategory)

	(*router).Delete("/:id", controller.DeleteCategoryById)

	(*router).Post("", controller.CreateCategory)

	(*router).Put("", controller.UpdateCategory)
}
func Setup(app *fiber.App) {
	app.Get("/image", controller.GetAllImageProduct)
	app.Delete("/image/:id", controller.DelReviewByID)
	app.Post("/image", controller.CreateImageProduct)

	//review
	app.Get("/review", controller.GetAllReview)
	app.Delete("/review/:id", controller.DelReviewByID)
	app.Post("/review", controller.CreateReview)
	//History
	app.Get("/history", controller.GetAllHistoryPrices)
}
