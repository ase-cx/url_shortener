package handlers

import (
	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	shorten := c.Params("s")

	var url models.URL
	err := database.DB.Collection("urls").FindOne(c.Context(), fiber.Map{"shorten": shorten}).Decode(&url)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cannot find URL"})
	}
	// Redirect to original URL
	return c.Redirect(url.Original, fiber.StatusFound)
}
