package main

import (
	"log"

	"github.com/Justincletus003/go-prod-ord/database"
	"github.com/Justincletus003/go-prod-ord/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func setupRoutes(app *fiber.App)  {
	app.Get("/api", welcome)
	// User details
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Product details
	app.Get("/api/products", routes.GetProducts)
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// order details
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Put("/api/orders/:id", routes.UpdateOrder)
	app.Delete("/api/orders/:id", routes.DeleteOrder)

}


func main() {
	database.Connect()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
