package middleware

import (
	"context"
	pk "github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/gofiber/fiber/v2"
	"time"
)

func AdminMiddleware(timeout time.Duration, client pk.AuthClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, ok := c.Locals("id").(int64)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"message": "неавторизированный пользователь",
			})
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		req := &pk.IsAdminRequest{
			UserId: id,
		}

		res, err := client.IsAdmin(ctx, req)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if !res.IsAdmin {
			return c.Status(401).JSON(fiber.Map{
				"message": "отказ в доступе",
			})
		}

		return c.Next()
	}
}
