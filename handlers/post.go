// @/handlers/posts.go
package handlers

import (
	"back/config"
	"back/entities"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []entities.Post

	config.Database.Find(&posts)
	return c.Status(200).JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entities.Post

	result := config.Database.Find(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&post)
}

func AddPost(c *fiber.Ctx) error {
	post := new(entities.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&post)
	return c.Status(201).JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	post := new(entities.Post)
	id := c.Params("id")

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&post)
	return c.Status(200).JSON(post)
}

func RemovePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post entities.Post

	result := config.Database.Delete(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
