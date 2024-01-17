package handlers

import (
	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

func Shorten(c *fiber.Ctx) error {
	url := new(models.URL)
	if err := c.BodyParser(url); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// if shorten is empty, generate new shorten
	// TODO: Implement this for now just return error
	if url.Shorten == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten cannot be empty"})
	}

	// Check if shorten is already exist
	var foundURL models.URL
	err := database.DB.Collection("urls").FindOne(c.Context(), fiber.Map{"shorten": url.Shorten}).Decode(&foundURL)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten already exist"})
	}

	// Insert url into database
	_, err = database.DB.Collection("urls").InsertOne(c.Context(), url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.JSON(url)
}
