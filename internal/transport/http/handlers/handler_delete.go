package handlers

import (
	"errors"
	pk "github.com/DrusGalkin/auth-protos/gen/go/auth"
	"github.com/DrusGalkin/user-info/internal/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/context"
)

func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Невалидный id",
		})
	}

	req := &pk.IsAdminRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	req.UserId = int64(id)

	res, err := h.Auth.IsAdmin(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res.IsAdmin {
		return c.Status(500).JSON(fiber.Map{
			"error": "Вы не можете удалить админа",
		})
	}

	err = h.uc.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return c.Status(201).JSON(fiber.Map{
				"message": "пользователя не существует",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "Ошибка удаления пользователя",
		})
	}

	return c.Status(200).JSON(true)
}

func (h *UserHandler) DeleteUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	req := &pk.IsAdminRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	user, err := h.uc.UserByUsername(username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req.UserId = int64(user.ID)

	res, err := h.Auth.IsAdmin(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res.IsAdmin {
		return c.Status(500).JSON(fiber.Map{
			"error": "Вы не можете удалить админа",
		})
	}

	err = h.uc.DeleteUserByUsername(username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(true)
}

func (h *UserHandler) DeleteUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	req := &pk.IsAdminRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), h.Timeout)
	defer cancel()

	user, err := h.uc.UserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req.UserId = int64(user.ID)

	res, err := h.Auth.IsAdmin(ctx, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if res.IsAdmin {
		return c.Status(500).JSON(fiber.Map{
			"error": "Вы не можете удалить админа",
		})
	}

	err = h.uc.DeleteUserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(true)
}
