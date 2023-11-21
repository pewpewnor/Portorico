package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/src/model"
)

func (h *Handler) CreateWebsite(c *fiber.Ctx) error {
	var body struct {
		Name         string `json:"name"`
		TemplateName string `json:"templateName"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	user, _ := c.Locals("user").(model.User)

	website, err := h.websiteRepository.CreateWebsite(body.Name, body.TemplateName, user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(website)
}
