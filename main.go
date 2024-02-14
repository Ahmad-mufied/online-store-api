package main

import (
	"fmt"
	"github.com/Ahmad-mufied/online-store/database"
	"github.com/Ahmad-mufied/online-store/router"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})

	database.ConnectDB()

	r := router.SetupRoutes(app)
	fmt.Println(r)

	log.Fatal(app.Listen(":8080"))
}
