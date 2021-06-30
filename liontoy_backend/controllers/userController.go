package controllers

import (
	"strconv"

	"github.com/TechMaster/golang/08Fiber/Basic/database"
	"github.com/TechMaster/golang/08Fiber/Basic/models"
	"github.com/gofiber/fiber/v2"
)

func AllUser(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")
	database.DB.Create(&user)
	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.DB.Find(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	// if err := c.BodyParser(&user); err != nil {
	// 	return err
	// }
	database.DB.Find(&user)
	user.Name = "tung"
	user.Email = "dss"
	return c.JSON(user)
}
