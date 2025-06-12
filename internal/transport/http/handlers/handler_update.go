package handlers

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) UpdatePassword(c *fiber.Ctx) error {
	id := c.Locals("id")
	var req struct {
		id       int    `json:"id"`
		password string `json:"password"`
	}

	uid := id.(int)
	if uid <= 0 || id == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "пользователь не авторизирован",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "невалидные данные",
		})
	}

	req.id = uid

	err := h.uc.UpdatePassword(req.id, req.password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "пароль изменен",
	})
}
