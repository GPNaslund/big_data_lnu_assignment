package dataservice

import (
	"1dv027/wt2/internal/dto"
	"1dv027/wt2/internal/model"
	"context"
)

type QueryParamValidator interface {
	Validate(queryParams dto.DataQuery) (dto.DataRequest, error)
}

type VideoGamesRepository interface {
	GetAggregatedData(ctx context.Context, request dto.DataRequest) ([]model.Dataset, error)
}

// Service that gets video games data.
type VideoGamesDataService struct {
	paramValidator QueryParamValidator
	repo           VideoGamesRepository
}

// Creates a new instance of VideoGamesDataService.
func NewVideoGamesDataService(paramValidator QueryParamValidator, repo VideoGamesRepository) VideoGamesDataService {
	return VideoGamesDataService{
		paramValidator: paramValidator,
		repo:           repo,
	}
}

// Method for getting video games data from repository.
func (d VideoGamesDataService) GetData(ctx context.Context, queryParams dto.DataQuery) ([]dto.Dataset, error) {
	dataRequest, err := d.paramValidator.Validate(queryParams)
	if err != nil {
		return nil, err
	}
	dataSlice, err := d.repo.GetAggregatedData(ctx, dataRequest)
	if err != nil {
		return nil, err
	}

	var result []dto.Dataset
	for _, dataset := range dataSlice {
		result = append(result, dto.Dataset{Data: dataset.Data})
	}
	return result, nil
}
