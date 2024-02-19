package handlers

import (
	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/models"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	contentType := c.Get("Content-Type")
	if contentType == "application/json" {
		// Parse body as JSON
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		// Parse form values
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported Content-Type"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	user.Password = string(hashedPassword)

	// Insert user into database
	_, err = database.DB.Collection("users").InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)

	contentType := c.Get("Content-Type")
	if contentType == "application/json" {
		// Parse body as JSON
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
	} else if contentType == "application/x-www-form-urlencoded" {
		// Parse form values
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unsupported Content-Type"})
	}

	// Find user in database declared placeholder as User type
	var foundUser models.User
	err := database.DB.Collection("users").FindOne(c.Context(), fiber.Map{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User not found"})
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// If the password is correct start J3K
	token := jwt.New(jwt.SigningMethodHS256)

	// Create claims which is like a mini database that stores jwt data
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = foundUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 1 day

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("THIS IS A SECRET")) // DO NOT USE THIS IN PRODUCTION
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create JWT Token for some reason"})
	}

	// Return the token
	return c.JSON(fiber.Map{"token": t})
}
