package videogamessearchhandler

import (
	"1dv027/wt2/internal/dto"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Service for getting search based data.
type GetVideoGamesSearchService interface {
	GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error)
}

// Handler for GET /video-games/search
type GetVideoGamesSearchHandler struct {
	service GetVideoGamesSearchService
}

// Creates a new GetVideoGamesSearchHandler
func NewGetVideoGamesSearchHandler(service GetVideoGamesSearchService) GetVideoGamesSearchHandler {
	return GetVideoGamesSearchHandler{
		service: service,
	}
}

// Method for handling incoming requests.
func (g GetVideoGamesSearchHandler) Handle(c *fiber.Ctx) error {
	searchParam := c.Query("query", "")
	if searchParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no search param provided",
		})
	}

	pageParam := c.Query("page", "")
	if pageParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no page param provided",
		})
	}

	pageParamInt, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "page parameter must be a number",
		})
	}

	ctx := c.Context()
	data, err := g.service.GetSearchData(ctx, searchParam, pageParamInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
