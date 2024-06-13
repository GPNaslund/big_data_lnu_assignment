package videogamesdatahandler

import (
	customerror "1dv027/wt2/internal/custom-errors"
	"1dv027/wt2/internal/dto"
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Service for getting video-games data
type GetDataService interface {
	GetData(ctx context.Context, queryParams dto.DataQuery) ([]dto.Dataset, error)
}

// Handler for GET /video-games/data
type GetVideoGamesDataHandler struct {
	service GetDataService
}

// Creates a new instance of GetVideoGamesDataHandler
func NewGetVideoGamesDataHandler(service GetDataService) GetVideoGamesDataHandler {
	return GetVideoGamesDataHandler{
		service: service,
	}
}

// Method for handling incoming requests.
func (g GetVideoGamesDataHandler) Handle(c *fiber.Ctx) error {
	queryParams := new(dto.DataQuery)
	err := c.QueryParser(queryParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not parse query parameters",
		})
	}

	ctx := c.Context()
	data, err := g.service.GetData(ctx, *queryParams)
	if err != nil {
		var queryParamErr *customerror.QueryParamError
		if errors.As(err, &queryParamErr) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid query parameter(s)",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong internally",
		})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}
