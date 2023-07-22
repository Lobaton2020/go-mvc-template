package main

import (
	"os"

	"github.com/Lobaton2020/go-mvc-template/initializers"
	"github.com/Lobaton2020/go-mvc-template/middleware"
	"github.com/Lobaton2020/go-mvc-template/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvironment()
	initializers.ConnectDatabase()
	initializers.SyncDB()
}
func main() {
	//templates set
	engine := html.New("./views", ".html")
	//server start
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//static
	app.Static("/", "./public")
	app.Use(middleware.Auth)
	//routes
	app.Get("/",func(c *fiber.Ctx) error {
		return c.Redirect("posts")
	})
	routes.RoutePosts(app)

	app.Listen(":" + os.Getenv("PORT"))
}
