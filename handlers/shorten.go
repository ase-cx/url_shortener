package handlers

import (
	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

func Shorten(c *fiber.Ctx) error {
    url := new(models.URL)

    contentType := c.Get("Content-Type")
    if contentType == "application/json" {
        // Parse body as JSON
        if err := c.BodyParser(url); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
        }
    } else if contentType == "application/x-www-form-urlencoded" {
        // Parse form values
        url.Original = c.FormValue("original")
        url.Shorten = c.FormValue("shorten")
        // Additional form value parsing can go here
    } else {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported Content-Type"})
    }

    // Continue with the rest of the function...
    if url.Shorten == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten cannot be empty"})
    }

    var foundURL models.URL
    err := database.DB.Collection("urls").FindOne(c.Context(), fiber.Map{"shorten": url.Shorten}).Decode(&foundURL)
    if err == nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Shorten already exists"})
    }

    _, err = database.DB.Collection("urls").InsertOne(c.Context(), url)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(url)
}
