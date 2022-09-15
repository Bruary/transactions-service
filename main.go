package main

// Requirements:
// 1) Add new transaction
// 2) Sum transactions
// 3) set/update budget
// 4) create new user
// 5)
import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// TODO:
	//1) Add transactions models
	//2) Integrate protobuffs
	v1.Post("/transactions", func(c *fiber.Ctx) {

	})

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(3000)
}
