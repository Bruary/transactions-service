package main

// Requirements:
// 1) Add new transaction
// 2) Sum transactions
// 3) set/update budget
// 4) create new user
// 5)
import (
	"encoding/json"
	"fmt"

	pb "github.com/Bruary/transactions-service/models/proto"
	"github.com/Bruary/transactions-service/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// TODO:
	//1) Add transactions models
	//2) Integrate protobuffs

	v1.Post("/transactions", func(c *fiber.Ctx) error {

		req := pb.CreateTransactionRequest{}

		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}

		fmt.Println("Coming here?")

		result, err := service.CreateTransaction(req)
		if err != nil {
			return err
		}

		resp, err := json.Marshal(&result)
		if err != nil {
			return err
		}

		c.JSON(resp)

		return nil
	})

	app.Listen(":3000")
}
