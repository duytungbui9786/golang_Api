package controller

import (
	"fmt"
	"time"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(repo.Product.GetAllProducts())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Product, err := repo.Product.FindProductById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(Product)
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.Product.DeleteProductById(int64(id))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	} else {
		return c.SendString("del Product successfully")
	}
}

func CreateProduct(c *fiber.Ctx) error {
	Product := new(model.Product)
	err := c.BodyParser(&Product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	ProductId := repo.Product.CreateNewProduct(Product)
	return c.SendString(fmt.Sprintf("New Product is created successfully with id = %d", ProductId))
}

func UpdateProduct(c *fiber.Ctx) error {
	updatedProduct := new(model.Product)

	err := c.BodyParser(&updatedProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	ProductPrice := repo.Product.GetPriceProductById(updatedProduct.Id)
	if updatedProduct.Price != ProductPrice {
		t := time.Now().String()
		HistoryPrice := new(model.HistoryPrice)
		HistoryPrice.IdProduct = updatedProduct.Id
		HistoryPrice.OldPrice = updatedProduct.Price
		HistoryPrice.Time = t
		err := c.BodyParser(&HistoryPrice)
		// if error
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Cannot parse JSON",
				"error":   err,
			})
		}
		errr := repo.HistoryPrice.CreateNewHistoryPrice(HistoryPrice)
		if errr != nil {
			return c.Status(404).SendString(err.Error())
		}
	}
	err = repo.Product.UpdateProduct(updatedProduct)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated ", updatedProduct.Id))
}

//History
func GetAllHistoryPrices(c *fiber.Ctx) error {
	return c.JSON(repo.HistoryPrice.GetAllHistoryPrices())
}
