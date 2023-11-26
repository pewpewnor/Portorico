package handlers

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pewpewnor/portorico/server/src/model"
)

func (h *handler) GetWebsite(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		return c.SendStatus(400)
	}

	website, found := h.websiteRepo.GetByName(name)
	if !found {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(website)
}

func (h *handler) GetWebsiteForEditing(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		return c.SendStatus(400)
	}

	website, found := h.websiteRepo.GetByName(name)
	if !found {
		return c.SendStatus(404)
	}

	user, _ := c.Locals("user").(model.User)
	if website.UserId != user.Id {
		return c.SendStatus(401)
	}

	return c.Status(200).JSON(website)
}

func (h *handler) FindWebsitesOwnedByUser(c *fiber.Ctx) error {
	user, _ := c.Locals("user").(model.User)
	websites, err := h.websiteRepo.FindByUserId(user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(websites)
}

func (h *handler) CreateWebsite(c *fiber.Ctx) error {
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
	h.validateStringMaxLength(validations, "templateName", "template name", 255, body.TemplateName)
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
	if _, exist := h.websiteRepo.GetByName(body.Name); exist {
		validations["name"] = "website name is already taken, please try a different one"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	user, _ := c.Locals("user").(model.User)
	website, err := h.websiteRepo.Create(body.Name, body.TemplateName, body.Description, user.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(map[string]any{"website": website})
}

func (h *handler) UpdateWebsiteContent(c *fiber.Ctx) error {
	var body struct {
		Content   string `json:"content"`
		WebsiteId string `json:"websiteId"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateJSONString(validations, "content", body.Content)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	websiteId, err := uuid.Parse(body.WebsiteId)
	if err != nil {
		validations["websiteId"] = "websiteId is invalid"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	website, found := h.websiteRepo.GetById(websiteId)
	if !found {
		return c.SendStatus(404)
	}

	user, _ := c.Locals("user").(model.User)
	if website.UserId != user.Id {
		return c.SendStatus(403)
	}

	err = h.websiteRepo.Update(websiteId, website.Name, website.Description, json.RawMessage(body.Content))
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}

func (h *handler) UpdateWebsiteInformation(c *fiber.Ctx) error {
	var body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		WebsiteId   string `json:"websiteId"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(400)
	}

	validations := map[string]string{}
	h.validateStringMaxLength(validations, "name", "website name", 64, body.Name)
	h.validateStringMaxLength(validations, "description", "description", 170, body.Description)
	if len(validations) > 0 {
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.Contains(body.Name, " ") {
		validations["name"] = "website name must not contain any spaces"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	if strings.ContainsAny(body.Name, "/&?=:%\\") {
		validations["name"] =
			"website name must not contain characters such as '/', '&', '?', '=', ':', '%', '\\'"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	websiteId, err := uuid.Parse(body.WebsiteId)
	if err != nil {
		validations["websiteId"] = "websiteId is invalid"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}
	website, found := h.websiteRepo.GetById(websiteId)
	if !found {
		return c.SendStatus(404)
	}
	if otherWebsite, exist := h.websiteRepo.GetByName(body.Name); exist && otherWebsite.Id != websiteId {
		validations["name"] = "website name is already taken, please try a different one"
		return c.Status(400).JSON(map[string]any{"validations": validations})
	}

	user, _ := c.Locals("user").(model.User)
	if website.UserId != user.Id {
		return c.SendStatus(403)
	}

	err = h.websiteRepo.Update(websiteId, body.Name, body.Description, website.Content)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}

func (h *handler) DeleteWebsite(c *fiber.Ctx) error {
	websiteId, err := uuid.Parse(c.Params("websiteId"))
	if err != nil {
		return c.SendStatus(400)
	}
	website, found := h.websiteRepo.GetById(websiteId)
	if !found {
		return c.SendStatus(404)
	}

	user, _ := c.Locals("user").(model.User)
	if website.UserId != user.Id {
		return c.SendStatus(403)
	}

	err = h.websiteRepo.Delete(websiteId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
