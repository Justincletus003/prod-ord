package routes

import (
	"errors"

	"github.com/Justincletus003/go-prod-ord/database"
	"github.com/Justincletus003/go-prod-ord/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`	
}

func ResponseUser(userModel models.User) User  {
	return User{Id: userModel.Id, FirstName: userModel.FirstName, LastName: userModel.LastName}	
}

func CreateUser(c *fiber.Ctx) error  {
	var user models.User
	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&user)
	responseUser := ResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error  {
	var users []models.User
	database.DB.Find(&users)
	// var responseUsers []users{}
	responseUsers := []User{}
	for _, user := range users{
		responseUser := ResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func FindUser(id int, user *models.User) error {
	database.DB.Find(&user, "id=?", id)
	if user.Id == 0 {
		return errors.New("User does not found!")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil{
		return c.Status(400).JSON("please send user :id number")
	}
	var user models.User
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := ResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please send correct user id")
	}
	var user models.User
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type Updates struct{
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
	var updateData Updates
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	database.DB.Save(&user)
	responseUser := ResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

func DeleteUser(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please send proper user id")
	}
	
	var user models.User
	if err := FindUser(id, &user); err != nil{
		return c.Status(400).JSON(err.Error())
	}
	if err := database.DB.Delete(&user).Error; err != nil{
		return c.Status(404).JSON(err.Error())
	}
	
	return c.Status(200).SendString("User deleted successfully!")

}

