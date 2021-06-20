package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReview())
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	err = repo.Product.CheckProduct(review)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	reviewID := repo.Reviews.CreateNewReview(review)
	product, err := repo.Product.FindProductById(int64(review.ProductId))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	result := repo.Reviews.AverageRating(int64(review.ProductId))
	product.Rating = float32(result[int64(review.ProductId)])
	// newRate := repo.Books.GetRateBook(review)
	return c.SendString(fmt.Sprintf("New reivew is created successfully with id = %d", reviewID))
}

func AverageRating(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	_, err := repo.Product.FindProductById(int64(id))

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Not found Product for this id",
		})
	}
	result := repo.Reviews.AverageRating(int64(id))
	return c.JSON(result[int64(id)])

}

func DelReviewByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Reviews.DeleteReviewById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	product, err := repo.Product.FindProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	result := repo.Reviews.AverageRating(int64(id))
	product.Rating = float32(result[int64(id)])
	return c.SendString("delete successfully")

}
