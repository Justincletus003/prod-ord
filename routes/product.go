package routes

import (
	"errors"

	"github.com/Justincletus003/go-prod-ord/database"
	"github.com/Justincletus003/go-prod-ord/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct{
	Id uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func ResponseProduct(modelProduct models.Product) Product  {
	return Product{Id: modelProduct.Id, Name: modelProduct.Name, SerialNumber: modelProduct.SerialNumber}
}


func GetProducts(c *fiber.Ctx) error  {
	var products []models.Product

	database.DB.Find(&products)
	responseProducts := []Product{}

	for _, product := range products {
		responseProduct := ResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func CreateProduct(c *fiber.Ctx) error  {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&product)
	responseProduct := ResponseProduct(product)

	return c.Status(201).JSON(responseProduct)
}

func FindProduct(id int, productModel *models.Product) error  {	
	database.DB.Find(&productModel, "id=?", id)
	if productModel.Id == 0 {
		return errors.New("product not found!")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please send proper product id")
	}
	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseProduct := ResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please send proper product id")
	}
	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	type Product struct{
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var productData Product

	if err := c.BodyParser(&productData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = productData.Name
	product.SerialNumber = productData.SerialNumber

	database.DB.Save(&product)

	responseProduct := ResponseProduct(product)
	return c.Status(200).JSON(responseProduct)	

}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(404).JSON("Please send proper product id")
	}
	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).SendString("product deleted successfully!")
}