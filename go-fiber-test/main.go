package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New() // * ใช้ libary (fiber)

	/*
		! Fiber.Ctx = Context
		* เช่น BodyParser

		BodyParser = การรับค่าข้อมูลมาจากทาง KeyBorad
	*/

	//* วิธีสร้าง Group API
	api := app.Group("/api")
	// v1
	v1 := api.Group("/v1")
	// ! Result : /api/v1

	// * [Middleware && Basic Authentication]
	// Provide a minimal config
	// ! ต้องไว้ด้านบนเสมอ
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",    // * username && password
			"admin": "123456", // * username && password
		},
	}))

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") //* แสดง Hello World !
	})

	// * เส้น POST (ส่งข้อมูล) [BodyParser : การรับข้อมูลในรูปแบบ JSON]
	type Person struct {
		//ส่วนเก็บข้อมูลไว้ใน GO
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	v1.Post("/", func(c *fiber.Ctx) error {
		p := new(Person) // เก็บข้อมูล

		if err := c.BodyParser(p); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("error")
			// return err
		}

		log.Println(p.Name) // john
		log.Println(p.Pass) // doe
		str := p.Name + p.Pass
		return c.JSON(str)
	})

	// * [Params]
	v1.Get("/user/:name", func(c *fiber.Ctx) error {

		str := "hello ==> " + c.Params("name")
		return c.JSON(str)
	})

	// * [Query]
	v1.Post("/inet", func(c *fiber.Ctx) error {
		a := c.Query("search") // ถ้า Search มาจาก หน้าบ้าน จะเก็บใน a
		str := "my search is  " + a
		return c.JSON(str)
	})

	//* [Validation]
	v1.Post("/valid", func(c *fiber.Ctx) error {
		//*Connect to database
		type User struct {
			Name     string `json:"name" validate:"required,min=3,max=32"`
			IsActive *bool  `json:"isactive" validate:"required"`
			Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
		}
		user := new(User)
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
	})

	/* (Status Code)
	! 400 : Fail
	! 401 : Token is invalid
	! 402 : Login is true but you not use this page (Access denied)
	! 403 : Record not found
	! 404 : Not found

	* 200 : Success
	* 201 : Create Success

	* 500 :  Internal Server Error
	* 502 : Sever Fail
	*/

	// * กำหนด Post (เปิด Server)
	app.Listen(":3000")
}
