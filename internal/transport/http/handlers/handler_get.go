package handlers

import "github.com/gofiber/fiber/v2"

func (h *UserHandler) UserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Невалидный id",
		})
	}

	user, err := h.uc.UserByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Locals("response", user)
	return c.Status(200).JSON(user)
}

func (h *UserHandler) UserByEmail(c *fiber.Ctx) error {
	user, err := h.uc.UserByEmail(c.Params("email"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Locals("response", user)
	return c.Status(200).JSON(user)
}

func (h *UserHandler) UserByUsername(c *fiber.Ctx) error {
	user, err := h.uc.UserByUsername(c.Params("username"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Locals("response", user)
	return c.Status(200).JSON(user)
}

func (h *UserHandler) AllUsers(c *fiber.Ctx) error {
	users, err := h.uc.AllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Locals("response", users)
	return c.Status(200).JSON(users)
}
