package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pewpewnor/portorico/server/src/model"
)

func (h *Handler) FindWebsitesOwnedByUser(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*model.User)
	if !ok {
		return c.SendStatus(500)
	}

	websites, err := h.websiteRepository.FindByUserId(user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(websites)
}

func (h *Handler) CreateWebsite(c *fiber.Ctx) error {
	var body struct {
		Name         string `json:"name"`
		TemplateName string `json:"templateName"`
		Description  string `json:"description"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateStringMaxLength(validations, "name", "name", 64, body.Name)
	h.validateStringMaxLength(validations, "templateName", "template name", 255, body.Name)
	h.validateStringMaxLength(validations, "description", "description", 170, body.Description)

	if strings.ContainsAny(body.Name, "/&?=:%\\") {
		validations["templateName"] = "template name must not contain characters such as '/', '&', '?', '=', ':', '%', '\\'"
	}
	if strings.Contains(body.Name, " ") {
		validations["templateName"] = "template name must not contain any spaces"
	}

	user, _ := c.Locals("user").(model.User)

	website, err := h.websiteRepository.Create(body.Name, body.TemplateName, body.Description, user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(website)
}
