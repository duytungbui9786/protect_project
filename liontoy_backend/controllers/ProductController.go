package controllers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/TechMaster/golang/08Fiber/Basic/database"
	"github.com/TechMaster/golang/08Fiber/Basic/models"
	repo "github.com/TechMaster/golang/08Fiber/Basic/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {

	var product []models.Product
	database.DB.Find(&product)
	return c.JSON(product)
}
func GetAllProductLimit(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 9
	offset := (page - 1) * limit
	var total int64
	var product []models.Product
	database.DB.Offset(offset).Limit(limit).Find(&product)
	database.DB.Model(&models.Product{}).Count(&total)
	return c.JSON(fiber.Map{
		"data": product,
		"meta": fiber.Map{
			"total":    total,
			"page":     page,
			"lastpage": math.Ceil(float64(int(total)/limit) + 1),
		},
	})
}
func GetAllProductBoy(c *fiber.Ctx) error {

	var product []models.Product
	// database.DB.Where("Agesrank =?", 1).Find(&product)
	// return c.JSON(product)
	database.DB.Where("Agesrank = ?", 1).Find(&product)
	return c.JSON(product)
}

func GetAllProductGirl(c *fiber.Ctx) error {

	var product []models.Product
	database.DB.Where("Agesrank = ?", 2).Find(&product)
	return c.JSON(product)

}

func Get5ProductRelate(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.Raw("SELECT * FROM `products` ORDER BY RAND() LIMIT 4").Scan(&product)
	return c.JSON(product)
}

func GetAllProductHot(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.Raw("SELECT * FROM `products` WHERE `Comming` = 1 ORDER BY RAND() LIMIT 8").Scan(&product)
	return c.JSON(product)
}

func GetAllProductFeature(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.Raw("SELECT * FROM `products` WHERE `Featured` = 1 ORDER BY RAND() LIMIT 8").Scan(&product)
	return c.JSON(product)
}

func GetAllProductNew(c *fiber.Ctx) error {
	var product []models.Product
	database.DB.Raw("SELECT * FROM `products` WHERE `New` =  1 ORDER BY RAND() LIMIT 8").Scan(&product)
	return c.JSON(product)
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
	Product := new(models.Product)
	err := c.BodyParser(&Product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	database.DB.Create(&Product)
	return c.JSON(Product)
	// err = repo.Product.CreateNewProduct(Product)
	return c.SendString(fmt.Sprintf("New Product is created successfully "))
}

func UpdateProduct(c *fiber.Ctx) error {
	updatedProduct := new(models.Product)

	err := c.BodyParser(&updatedProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.Product.UpdateProduct(updatedProduct)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated ", updatedProduct.Id))
}

//History
