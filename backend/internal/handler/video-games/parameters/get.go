package videogamesparameterhandler

import (
	"1dv027/wt2/internal/dto"
	"context"

	"github.com/gofiber/fiber/v2"
)

type GetVideoGamesParametersService interface {
	GetParameters(ctx context.Context) (dto.VideoGamesParameters, error)
}

// Handler for GET /video-games/parameters.
type GetVideoGamesParametersHandler struct {
	service GetVideoGamesParametersService
}

// Creates a new instance of GetVideoGamesParameterHandler.
func NewGetVideoGamesParametersHandler(service GetVideoGamesParametersService) GetVideoGamesParametersHandler {
	return GetVideoGamesParametersHandler{
		service: service,
	}
}

// Method for handling incoming request, will respond with all valid
// data parameters that can be used to query the database with.
func (g GetVideoGamesParametersHandler) Handle(c *fiber.Ctx) error {
	ctx := c.Context()
	parameters, err := g.service.GetParameters(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong internally",
		})
	}

	return c.Status(fiber.StatusOK).JSON(parameters)
}
