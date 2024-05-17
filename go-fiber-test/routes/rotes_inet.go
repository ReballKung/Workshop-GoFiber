package routes

import (
	"go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	/*
		! Fiber.Ctx = Context
		* เช่น BodyParser
		BodyParser = การรับค่าข้อมูลมาจากทาง KeyBorad
	*/

	//* วิธีสร้าง Group API
	api := app.Group("/api")
	// v1
	v1 := api.Group("/v1")
	// v2
	v2 := api.Group("/v2")
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

	v1.Get("/", controllers.HelloTest) // * เรียกใช้ function จาก Controllers

	v2.Get("/", controllers.HelloTestV2) // * เรียกใช้ function จาก Controllers

	// * เส้น POST (ส่งข้อมูล) [BodyParser : การรับข้อมูลในรูปแบบ JSON]
	v1.Post("/", controllers.BodyParserTest)

	// * [Params]
	v1.Get("/user/:name", controllers.ParamsTest)

	// * [Query]
	v1.Post("/inet", controllers.QueryTest)

	//* [Validation]
	v1.Post("/valid", controllers.ValidTest)

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
}
