package parametersservice

import (
	"1dv027/wt2/internal/dto"
	"context"
)

type ValidParameters interface {
	GetAggregations() []string
	GetGroupings() []string
	GetRegions() []string
	GetGenres() []string
	GetStartYear() int
	GetEndYear() int
}

// Service for getting the valid parameters to use
// for querying the database for data.
type VideoGamesParametersService struct {
	params ValidParameters
}

// Creates a new instance of VideoGamesParametersService.
func NewVideoGamesParametersService(params ValidParameters) VideoGamesParametersService {
	return VideoGamesParametersService{
		params: params,
	}
}

// Method that returns valid parameters for data querying.
func (v VideoGamesParametersService) GetParameters(ctx context.Context) (dto.VideoGamesParameters, error) {
	validAggregations := v.params.GetAggregations()
	validGroups := v.params.GetGroupings()
	validRegions := v.params.GetRegions()
	validGenres := v.params.GetGenres()
	validStartYear := v.params.GetStartYear()
	validEndYear := v.params.GetEndYear()

	result := dto.VideoGamesParameters{
		Aggregations: validAggregations,
		Groupings:    validGroups,
		Regions:      validRegions,
		Genres:       validGenres,
		StartYear:    validStartYear,
		EndYear:      validEndYear,
	}

	return result, nil
}
