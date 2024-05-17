package controllers

import (
	"go-fiber-test/models"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!") //* แสดง Hello World !
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! V2") //* แสดง Hello World !
}
func BodyParserTest(c *fiber.Ctx) error {
	p := new(models.Person) // เก็บข้อมูล //* เรียกมาจาก models

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("error")
		// return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search") // ถ้า Search มาจาก หน้าบ้าน จะเก็บใน a
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//*Connect to database

	user := new(models.User) //* เรียกมาจาก models
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	// * nil = null

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	//  * แบบลดรูป
	// if errors := validate.Struct(user); errors != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	// }

	return c.JSON(user)
}
