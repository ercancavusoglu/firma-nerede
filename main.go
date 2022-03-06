package main

import (
	"github.com/ercancavusoglu/firma-nerede/database"
	"github.com/ercancavusoglu/firma-nerede/router"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/dashboard", monitor.New())
	app.Use(logger.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3001"))
}
