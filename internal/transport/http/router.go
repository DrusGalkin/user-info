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

	userCache := api.Group("/users").Use(middleware.CacheMiddleware(rdb))
	{
		userCache.Get("", handler.AllUsers)
		userCache.Get("/:id", handler.UserByID)
		userCache.Get("/username/:username", handler.UserByUsername)
		userCache.Get("/email/:email", handler.UserByEmail)
	}

	admin := api.Use(
		middleware.AuthMiddleware(timeout, client),
		middleware.AuthMiddleware(timeout, client),
	)
	{
		admin.Delete("/:id", handler.DeleteUserByID)
		admin.Delete("/username/:username", handler.DeleteUserByUsername)
		admin.Delete("/email/:email", handler.DeleteUserByEmail)
	}

	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware(timeout, client))
	{
		user.Patch("/password", handler.UpdatePassword)
	}

	return api
}
