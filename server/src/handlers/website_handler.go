package handlers

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	h.validateStringMaxLength(validations, "name", "website name", 64, body.Name)
	h.validateStringMaxLength(validations, "templateName", "template name", 255, body.Name)
	h.validateStringMaxLength(validations, "description", "description", 170, body.Description)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.Contains(body.Name, " ") {
		validations["name"] = "website name must not contain any spaces"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.ContainsAny(body.Name, "/&?=:%") {
		validations["name"] =
			"website name must not contain characters such as '/', '&', '?', '=', ':', '%'"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if h.websiteRepository.GetByName(body.Name) != nil {
		validations["name"] = "website name is already taken, please try a different one"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	user, _ := c.Locals("user").(*model.User)

	website, err := h.websiteRepository.Create(body.Name, body.TemplateName, body.Description, user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(website)
}

func (h *Handler) UpdateWebsite(c *fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Content     string `json:"content"`
		WebsiteId   string `json:"websiteId"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateStringMaxLength(validations, "name", "website name", 64, body.Name)
	h.validateStringMaxLength(validations, "templateName", "template name", 255, body.Name)
	h.validateStringMaxLength(validations, "description", "description", 170, body.Description)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.Contains(body.Name, " ") {
		validations["name"] = "website name must not contain any spaces"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.ContainsAny(body.Name, "/&?=:%") {
		validations["name"] =
			"website name must not contain characters such as '/', '&', '?', '=', ':', '%'"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if h.websiteRepository.GetByName(body.Name) != nil {
		validations["name"] = "website name is already taken, please try a different one"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	websiteId, err := uuid.Parse(body.WebsiteId)
	if err != nil {
		validations["websiteId"] = "websiteId is invalid"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	if h.websiteRepository.GetById(websiteId) == nil {
		return c.SendStatus(400)
	}
	err = h.websiteRepository.Update(body.Name, body.Description, json.RawMessage(body.Content), websiteId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
