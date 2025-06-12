package handlers

import (
	"github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Handler interface {
	UserByID(c *fiber.Ctx) error
	UserByEmail(c *fiber.Ctx) error
	UserByUsername(c *fiber.Ctx) error
	AllUsers(c *fiber.Ctx) error

	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error

	DeleteUserByID(c *fiber.Ctx) error
	DeleteUserByUsername(c *fiber.Ctx) error
	DeleteUserByEmail(c *fiber.Ctx) error

	UpdatePassword(c *fiber.Ctx) error
}

type UserHandler struct {
	uc      usecase.Usecase
	Timeout time.Duration
	Auth    auth.AuthClient
}

func New(uc usecase.Usecase, timeout time.Duration, client auth.AuthClient) Handler {
	return &UserHandler{
		uc:      uc,
		Timeout: timeout,
		Auth:    client,
	}
}
