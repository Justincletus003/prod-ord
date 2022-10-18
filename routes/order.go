package routes

import (
	"errors"
	"fmt"
	// "os"

	//"log"
	"time"

	"github.com/Justincletus003/go-prod-ord/database"
	"github.com/Justincletus003/go-prod-ord/models"

	// "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	Id        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func ResponseOrder(order models.Order, user User, product Product) Order {
	return Order{Id: order.Id, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	responseOrders := []Order{}

	database.DB.Find(&orders)

	for _, order := range orders{
		var user models.User
		var product models.Product
		database.DB.Find(&user, "id=?", order.UserRefer)
		database.DB.Find(&product, "id=?", order.ProductRefer)
		// if err := FindUser(order.UserRefer, &user); err != nil {
		// 	return c.Status(404).JSON(err.Error())
		// }
		// var product models.Product
		// if err := FindProduct(order.ProductRefer, &product); err != nil {
		// 	return c.Status(404).JSON(err.Error())
		// }
		responseUser := ResponseUser(user)
		responseProduct := ResponseProduct(product)
		responseOrder := ResponseOrder(order, responseUser, responseProduct)

		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&order)
	responseUser := ResponseUser(user)
	responseProduct := ResponseProduct(product)
	responseOrder := ResponseOrder(order, responseUser, responseProduct)
	return c.Status(201).JSON(responseOrder)

}

func FindOrder(id int, order *models.Order) error {
	database.DB.Find(&order, "id=?", id)
	if order.Id == 0 {
		return errors.New("Order not found!")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error  {
	var order models.Order
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON("Please send proper order id")
	}

	if err := FindOrder(id, &order); err != nil {
		c.Status(404).JSON(err.Error())
	}

	var user models.User
	database.DB.First(&user, "id=?", order.UserRefer)
	// if err:= FindUser(order.UserRefer, &user); err != nil {
	// 	return c.Status(404).JSON(err.Error())
	// }

	var product models.Product
	database.DB.First(&product, "id=?", order.ProductRefer)
	// if err := FindProduct(order.ProductRefer, &product); err != nil {
	// 	return c.Status(404).JSON(err.Error())
	// }

	responseUser := ResponseUser(user)
	responseProduct := ResponseProduct(product)
	responseOrder := ResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}

func UpdateOrder(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please send proper order id")
	}
	var order models.Order
	if err := FindOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	fmt.Println(order)

	type UpdateData struct {
		UserId uint `json:"user_id"`
		ProdId uint `json:"product_id"`
	}

	var orderData UpdateData
	
	if err := c.BodyParser(&orderData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	//log.Println(Order)
	var user models.User
	var product models.Product

	if err := FindUser(int(orderData.UserId), &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := FindProduct(int(orderData.ProdId), &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	order.User = user
	order.Product = product
	database.DB.Save(&order)
	responseOrder := ResponseOrder(order, ResponseUser(user), ResponseProduct(product))
	return c.Status(200).JSON(responseOrder)
}

func DeleteOrder(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please provide correct order id!")
	}
	var order models.Order

	if err := FindOrder(id, &order); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	if err := database.DB.Delete(&order).Error; err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("order deleted!")

}
