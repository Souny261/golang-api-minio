package main

// func main() {
// 	app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})
// 	log.Fatal(app.Listen(":3000"))
// }

import (
	"fmt"
	"go-fiber-minio/config"
	"go-fiber-minio/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Hello World",
		})
	})
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	_, err := minioUpload.MinioConnection()
	// 	if err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 			"error": true,
	// 			"msg":   err.Error(),
	// 		})
	// 	}
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"success": true,
	// 		"msg":     "OKOKOKOKO",
	// 	})
	// })
	app.Post("/upload", controller.UploadFile)
	app.Get("/getfile", controller.GetFile)
	app.Get("/getBytes", controller.GetFileBytes)
	MYPORT := config.GetEnv("app.port", "3000")
	SERVER_RUNNING := fmt.Sprintf(":%v", MYPORT)
	app.Listen(SERVER_RUNNING)
}
