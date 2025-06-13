package middleware

import (
	"context"
	"fmt"
	pk "github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/gofiber/fiber/v2"
	"time"
)

func AuthMiddleware(timeout time.Duration, client pk.AuthClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		req := &pk.ValidTokenRequest{
			Access: header,
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		claims, err := client.ValidToken(ctx, req)
		fmt.Println(claims)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "невалидный токен",
			})
		}

		c.Locals("id", fmt.Sprint(claims.UserId))
		c.Locals("email", claims.Email)
		return c.Next()
	}
}
