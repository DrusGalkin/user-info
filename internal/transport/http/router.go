package http

import (
	pk "github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/storage/redis"
	"github.com/DrusGalkin/user-info/internal/transport/http/handlers"
	"github.com/DrusGalkin/user-info/internal/transport/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)

func SetupRouters(handler handlers.Handler, timeout time.Duration, client pk.AuthClient, rdb *redis.App) *fiber.App {
	api := fiber.New()
	api.Use(logger.New())
	api.Use(cors.New())

	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)
	api.Post("/refresh", handler.Refresh)

	user := api.Group("/users")
	{
		user.Get("/:id", handler.UserByID)
		user.Get("/username/:username", handler.UserByUsername)
		user.Get("/email/:email", handler.UserByEmail)
		
		cacheUser := user.Use(middleware.CacheMiddleware(rdb))
		{
			cacheUser.Get("", handler.AllUsers)
		}
	}

	admin := user.Use(middleware.AuthMiddleware(timeout, client))
	{
		admin.Delete("/:id", handler.DeleteUserByID)
		admin.Delete("/username/:username", handler.DeleteUserByUsername)
		admin.Delete("/email/:email", handler.DeleteUserByEmail)
	}

	return api
}
