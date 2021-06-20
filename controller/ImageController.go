package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllImageProduct(c *fiber.Ctx) error {
	return c.JSON(repo.ImageProduct.GetAllImageProducts())
}

func DeleteImageProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = repo.ImageProduct.DeleteImageProductById(int64(id))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	} else {
		return c.SendString("del ImageProduct successfully")
	}
}

func CreateImageProduct(c *fiber.Ctx) error {
	ImageProduct := new(model.ImageProduct)
	err := c.BodyParser(&ImageProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	_, errr := repo.Product.FindProductById(int64(ImageProduct.IdProduct))
	if errr != nil {
		return c.Status(400).SendString("ERROR ! product not exists")
	}
	ImageProductId := repo.ImageProduct.CreateNewImageProduct(ImageProduct)
	return c.SendString(fmt.Sprintf("New ImageProduct is created successfully with id = %d", ImageProductId))
}
