package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	agent:=c.Get("User-Agent")
	fmt.Println(time.Local,agent)
	c.Set("X-Custom-Header", "Hello, World")
	return c.Next()
}