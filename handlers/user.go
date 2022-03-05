// @/handlers/user.go
package handlers

import (
	"back/config"
	"back/entities"

	"github.com/gofiber/fiber/v2"
)

// Tous les utilisateurs
func GetUsers(c *fiber.Ctx) error {
	var users []entities.User

	config.Database.Find(&users)
	return c.Status(200).JSON(users)
}

// Utilisateur spécifique par id
func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	var user entities.User
	result := config.Database.Find(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

// Modifications des infos d'un utilisateur
func UpdateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

// Suppression d'un utilisateur
func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entities.User

	result := config.Database.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// Inscription au site
func Signup(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	username := user.Username
	result := config.Database.Find(&user, "username = ?", username)
	if result.RowsAffected != 0 {
		return c.SendStatus(404)
	}

	config.Database.Create(&user)
	return c.Status(201).JSON(user)
}

// Connexion d'un utilisateur
func Login(c *fiber.Ctx) error {
	login := new(entities.Login)
	if err := c.BodyParser(login); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	username := login.Username

	var user entities.User
	result := config.Database.Find(&user, "username = ?", username)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	if login.Password == user.Password {
		return c.Status(200).JSON(&user)
	}

	return c.SendStatus(404)
}
