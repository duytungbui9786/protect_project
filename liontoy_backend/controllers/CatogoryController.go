package controllers

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Basic/models"
	repo "github.com/TechMaster/golang/08Fiber/Basic/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(c *fiber.Ctx) error {
	return c.JSON(repo.Category.GetAllCategorys())
}

func DeleteCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product := repo.Product.CheckProductUseCategory(int64(id))
	if product != nil {
		return c.Status(400).SendString("ERROR ! there is products using category")
	}
	err = repo.Category.DeleteCategoryById(int64(id))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	} else {
		return c.SendString("del Category successfully")
	}
}

func CreateCategory(c *fiber.Ctx) error {
	Category := new(models.Category)
	err := c.BodyParser(&Category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.SendString(fmt.Sprintf("New Category is created successfully with id = %d"))
}

func UpdateCategory(c *fiber.Ctx) error {
	updatedCategory := new(models.Category)

	err := c.BodyParser(&updatedCategory)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Category.UpdateCategory(updatedCategory)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Category with id = %d is successfully updated", updatedCategory.Id))
}
