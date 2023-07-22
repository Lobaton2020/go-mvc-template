package routes

import (
	"github.com/Lobaton2020/go-mvc-template/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutePosts(app *fiber.App) {
	postRouter := app.Group("posts")
	postController := controllers.Post{}
	postRouter.Get("/", postController.Index)
	postRouter.Get("/create", postController.Create)
}