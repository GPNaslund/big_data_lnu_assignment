package authmiddleware

import "github.com/gofiber/fiber/v2"

// Struct for validating api key.
type AuthMiddleware struct {
	apiKey string
}

// Creates a new auth middleware instance.
func NewAuthMiddleware(apiKey string) AuthMiddleware {
	return AuthMiddleware{
		apiKey: apiKey,
	}
}

// Handles incoming request, returns 401 on invalid api key.
func (a AuthMiddleware) Handle(c *fiber.Ctx) error {
	providedKey := c.Get("X-API-Key", "invalid")
	if providedKey == "invalid" || providedKey != a.apiKey {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid authentication",
		})
	}

	return c.Next()
}
