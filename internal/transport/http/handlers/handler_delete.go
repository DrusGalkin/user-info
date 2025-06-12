package handlers

import (
	"errors"
	"github.com/DrusGalkin/user-info/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Невалидный id",
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

	return c.Status(200).JSON(fiber.Map{
		"delete": true,
	})
}

func (h *UserHandler) DeleteUserByUsername(c *fiber.Ctx) error {
	err := h.uc.DeleteUserByUsername(c.Params("username"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"delete": true,
	})
}

func (h *UserHandler) DeleteUserByEmail(c *fiber.Ctx) error {
	err := h.uc.DeleteUserByEmail(c.Params("email"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"delete": true,
	})
}
