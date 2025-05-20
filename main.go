package main

import (
	app2 "intern_247/app"
	"intern_247/routes"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app2.Setup()
	app := fiber.New()
	routes.AdminRoutes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // fallback for local dev
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
