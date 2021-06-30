package controllers

import (
	"strconv"
	"time"

	"github.com/TechMaster/golang/08Fiber/Basic/database"
	"github.com/TechMaster/golang/08Fiber/Basic/models"
	"github.com/TechMaster/golang/08Fiber/Basic/util"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	if data["Password"] != data["Password_confim"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "wrong Password confim",
		})
	}

	user := models.User{
		Name:  data["Name"],
		Email: data["Email"],
		Quyen: data["Quyen"],
	}
	user.SetPassword(data["Password"])
	database.DB.Create(&user)
	return c.JSON(user)
}
func Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	var user models.User

	database.DB.Where("email", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "not found",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "wrong password",
		})
	}
	// tạo token đăng nhập
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	//đưa vào cookie
	//b1 : tọa cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	//bước 2 set cookie
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	//Check cookie
	cookie := c.Cookies("jwt")
	//get id form cookie
	id, _ := util.ParseJwt(cookie)
	//connect get user
	var user models.User
	database.DB.Where("id", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
