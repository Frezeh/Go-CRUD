package main

import (
	"fmt"
	"log"

	"github.com/Frezeh/golang/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func ping(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Get("/", ping)

	app.Get("/api/products", handlers.GetAllProducts)
	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products/:id", handlers.GetOneProduct)
	app.Patch("/api/products/:id", handlers.UpdateProduct)
	app.Delete("/api/products/:id", handlers.DeleteOneProduct)

	fmt.Println("Hello world")

	log.Fatal(app.Listen(":3000"))

}
