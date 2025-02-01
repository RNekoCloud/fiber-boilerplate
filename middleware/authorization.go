package middleware

import (
	"api-service/http"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	JwtSecret []byte
}

func (m *Middleware) Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: m.JwtSecret},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(401).JSON(&http.WebResponse{
				Status:  "error",
				Message: "Couldn't access resource because unauthorized request!",
				Data:    nil,
			})
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			c.Locals("name", claims["name"])
			c.Locals("username", claims["username"])
			c.Locals("user_id", claims["user_id"])
			c.Locals("role", claims["role"])

			return c.Next()
		},
	})
}
