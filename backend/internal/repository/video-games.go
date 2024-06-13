package repository

import (
	"1dv027/wt2/internal/dto"
	"1dv027/wt2/internal/model"
	"context"
	"encoding/json"
	"fmt"
)

type VideoGamesDataAccess interface {
	GetAggregateTotalByRegions(ctx context.Context) ([]dto.SalesData, error)
	GetAggregateTotalByGenres(ctx context.Context) ([]dto.SalesData, error)
	GetAggregateByYearByRegions(ctx context.Context, startYear int, endYear int, regions []model.ValidRegion) ([]dto.YearlySalesData, error)
	GetAggregateByYearByGenres(ctx context.Context, startYear int, endYear int, genres []model.ValidGenre) ([]dto.YearlySalesData, error)
	GetAggregateAllGames(ctx context.Context, startYear int, endYear int) ([]dto.VideoGame, error)
	GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error)
}

// Struct for simplifying data retrival.
type VideoGamesRepo struct {
	dataaccess VideoGamesDataAccess
}

// Creates a new instance of VideoGamesRepo.
func NewVideoGamesRepo(dataaccess VideoGamesDataAccess) VideoGamesRepo {
	return VideoGamesRepo{
		dataaccess: dataaccess,
	}
}

// Method for getting aggregated data based on the aggregation type in the data request.
func (v VideoGamesRepo) GetAggregatedData(ctx context.Context, request dto.DataRequest) ([]model.Dataset, error) {
	if request.Aggregation == model.AggregationTotal {
		return v.getAggregatedTotal(ctx, request)
	}

	if request.Aggregation == model.AggregationByYear {
		return v.getAggregatedByYear(ctx, request)
	}

	if request.Aggregation == model.AggregationAllGames {
		return v.getAggregatedAllGames(ctx, request)
	}

	return nil, fmt.Errorf("no valid data aggregate provided")
}

// Method for getting data based on search.
func (v VideoGamesRepo) GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error) {
	return v.dataaccess.GetSearchData(ctx, search, page)
}

// Gets aggregated sales data based on year.
func (v VideoGamesRepo) getAggregatedByYear(ctx context.Context, request dto.DataRequest) ([]model.Dataset, error) {
	startYear := request.StartYear
	endYear := request.EndYear

	if request.Group == model.GroupingByRegion {
		data, err := v.dataaccess.GetAggregateByYearByRegions(ctx, startYear, endYear, request.Regions)
		if err != nil {
			return nil, err
		}
		datasets, err := datasetsConverter(data)
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}

	if request.Group == model.GroupingByGenre {
		data, err := v.dataaccess.GetAggregateByYearByGenres(ctx, startYear, endYear, request.Genres)
		if err != nil {
			return nil, err
		}
		datasets, err := datasetsConverter(data)
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}
	return nil, fmt.Errorf("no valid group for aggregated by year provided")
}

// Gets total sales data aggregated by region or genre.
func (v VideoGamesRepo) getAggregatedTotal(ctx context.Context, request dto.DataRequest) ([]model.Dataset, error) {
	if request.Group == model.GroupingByRegion {
		data, err := v.dataaccess.GetAggregateTotalByRegions(ctx)
		if err != nil {
			return nil, err
		}
		datasets, err := datasetsConverter(data)
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}

	if request.Group == model.GroupingByGenre {
		data, err := v.dataaccess.GetAggregateTotalByGenres(ctx)
		if err != nil {
			return nil, err
		}
		datasets, err := datasetsConverter(data)
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}
	return nil, fmt.Errorf("no valid grouping for total aggregate provided")
}

// Gets all games aggregated by year.
func (v VideoGamesRepo) getAggregatedAllGames(ctx context.Context, request dto.DataRequest) ([]model.Dataset, error) {
	data, err := v.dataaccess.GetAggregateAllGames(ctx, request.StartYear, request.EndYear)
	if err != nil {
		return nil, err
	}
	datasets, err := datasetsConverter(data)
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

// Not tied to struct due to limitations of generics in Go.
func datasetsConverter[T any](dataSlice []T) ([]model.Dataset, error) {
	var datasets []model.Dataset
	for _, data := range dataSlice {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		datasets = append(datasets, model.Dataset{Data: jsonData})
	}
	return datasets, nil
}
