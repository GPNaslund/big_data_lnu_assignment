package searchservice

import (
	"1dv027/wt2/internal/dto"
	"context"
)

type VideoGamesSearchRepository interface {
	GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error)
}

// Service for querying the repository to get data based on search.
type VideoGamesSearchService struct {
	repo VideoGamesSearchRepository
}

// Creates a new instance of VideoGamesSearchService
func NewVideoGamesSearchService(repo VideoGamesSearchRepository) VideoGamesSearchService {
	return VideoGamesSearchService{
		repo: repo,
	}
}

// Method for getting search data from repository.
func (v VideoGamesSearchService) GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error) {
	return v.repo.GetSearchData(ctx, search, page)
}
