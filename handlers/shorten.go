package handlers

import (
	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

func Shorten(c *fiber.Ctx) error {
    url := new(models.URL)

    // Try to parse as JSON first
    jsonParseErr := c.BodyParser(url)
    
    // If JSON parsing fails, attempt to read from form data
    if jsonParseErr != nil {
        url.Original = c.FormValue("original")
        url.Shorten = c.FormValue("shorten")
        // Validate if required fields are populated
        if url.Original == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Original URL cannot be empty"})
        }
    }

    // if shorten is empty, generate new shorten
    // TODO: Implement this for now just return error
    if url.Shorten == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten cannot be empty"})
    }

    // Check if shorten already exists
    var foundURL models.URL
    err := database.DB.Collection("urls").FindOne(c.Context(), bson.M{"shorten": url.Shorten}).Decode(&foundURL)
    if err == nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten already exist"})
    }

    // Insert url into database
    _, err = database.DB.Collection("urls").InsertOne(c.Context(), url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(url)
}
