package controllers

import (
	"log"

	"github.com/Lobaton2020/go-mvc-template/initializers"
	"github.com/Lobaton2020/go-mvc-template/models"
	"github.com/gofiber/fiber/v2"
)

type Post struct{}

func (p *Post) Index(c *fiber.Ctx) error {
	var items = []models.Post{}
	initializers.DB.Find(&items)
	log.Print(items)
	return c.Render("posts/index", items)

}
func (p *Post) Create(c *fiber.Ctx) error {

	response := map[string]interface{}{
		"name":     "Andres",
		"lastname": "Lobaton",
		"age":      23,
	}
	return c.JSON(response)

}
