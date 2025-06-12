package middleware

import (
	"encoding/json"
	"github.com/DrusGalkin/user-info/internal/storage/redis"
	"github.com/gofiber/fiber/v2"
)

func CacheMiddleware(rdb *redis.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.OriginalURL()

		val, err := rdb.GetData(key)
		if err == nil {
			var data interface{}
			if err := json.Unmarshal([]byte(val), &data); err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "ошибка хеширования запроса",
				})
			}
			return c.JSON(data)
		}

		err = c.Next()

		if c.Response().StatusCode() == 200 {
			resData := c.Locals("response")
			resJSON, err := json.Marshal(resData)
			if err != nil {
				return err
			}
			if err := rdb.SetData(key, string(resJSON)); err != nil {
				return err
			}
		}
		return err
	}
}
