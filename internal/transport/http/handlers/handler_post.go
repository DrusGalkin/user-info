package handlers

import (
	"context"
	"errors"
	"github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	req := &auth.RegisterRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "невалидные данные",
		})
	}

	_, err := h.uc.UserByEmail(req.Email)
	if err != nil {
		if !errors.Is(err, repository.ErrUserNotFound) {
			return c.Status(400).JSON(fiber.Map{})
		}
	}

	id, err := h.Auth.Register(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(id)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	req := &auth.LoginRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "невалидные данные",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	tokens, err := h.Auth.Login(ctx, req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"access":  tokens.Access,
		"refresh": tokens.Refresh,
	})
}

func (h *UserHandler) Refresh(c *fiber.Ctx) error {
	req := &auth.RefreshRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "невалидный формат токена",
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()
	tokens, err := h.Auth.Refresh(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "ошибка генерации токена",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"access":  tokens.Access,
		"refresh": tokens.Refresh,
	})
}
