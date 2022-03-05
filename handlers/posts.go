// @/handlers/posts.go
package handlers

import (
	"back/config"
	"back/entities"

	_ "back/docs"

	"github.com/gofiber/fiber/v2"
)

// Tous les publications
func GetPosts(c *fiber.Ctx) error {
	var posts []entities.Post

	config.Database.Find(&posts)
	return c.Status(200).JSON(posts)
}

// Tous les publications d'un utilisateur
func GetPostsByUser(c *fiber.Ctx) error {
	var posts []entities.Post
	res := config.Database.Raw("SELECT * FROM posts WHERE username = ?", c.Params("user"))

	res.Scan(&posts)

	return c.Status(200).JSON(posts)
}

// Publication spécifique par id
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entities.Post

	result := config.Database.Find(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&post)
}

// Ajout d'une publication
func AddPost(c *fiber.Ctx) error {
	post := new(entities.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&post)
	return c.Status(201).JSON(post)
}

// Modification d'une publication
func UpdatePost(c *fiber.Ctx) error {
	post := new(entities.Post)
	id := c.Params("id")

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&post)
	return c.Status(200).JSON(post)
}

// Supprimer une publiation
func RemovePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entities.Post

	result := config.Database.Delete(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
