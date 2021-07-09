package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	// pusherClient := pusher.Client{
	// 	AppID: os.Getenv("PUSHER_APPID"),
	// 	Key: os.Getenv("PUSHER_KEY"),
	// 	Secret: os.Getenv("PUSHER_SECRET"),
	// 	Cluster: os.Getenv("PUSHER_CLUSTER"),
	// 	Secure: true,
	// }

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
