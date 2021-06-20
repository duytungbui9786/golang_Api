package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(repo.User.GetAllUsers())
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	User, err := repo.User.FindUserById(int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(User)
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = repo.User.DeleteUserById(int64(id))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	} else {
		return c.SendString("del user successfully")
	}
}

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	er := repo.User.CheckUserByUsername(user)
	if er != nil {
		return c.Status(404).SendString("user exist")
	}
	userId := repo.User.CreateNewUser(user)
	return c.SendString(fmt.Sprintf("New User is created successfully with id = %d", userId))
}

func UpdateUser(c *fiber.Ctx) error {
	updatedUser := new(model.User)

	err := c.BodyParser(&updatedUser)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.User.UpdateUser(updatedUser)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("User with id = %d is successfully updated", updatedUser.Id))
}
