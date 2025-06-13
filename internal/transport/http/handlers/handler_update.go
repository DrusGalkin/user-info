package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UpdateUser struct {
	NewPassword string `json:"password"`
}

func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	id := c.Locals("id")
	idStr := id.(string)
	uid, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "невалидный id пользователя",
		})
	}

	if uid <= 0 || id == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "пользователь не авторизирован",
		})
	}

	var u UpdateUser
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "невалидные данные",
		})
	}

	err = h.uc.UpdatePassword(uid, u.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "пароль изменен",
	})
}
