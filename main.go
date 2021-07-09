// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// )

// func main() {
// 	app := fiber.New()

// 	app.Use(cors.New())

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World 👋!")
// 	})

// 	app.Listen(":8000")
// }
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("SAMPLE"))
}
