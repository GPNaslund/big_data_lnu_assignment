package queryservice

import (
	customerror "1dv027/wt2/internal/custom-errors"
	"1dv027/wt2/internal/dto"
	"1dv027/wt2/internal/model"
	"strconv"
)

type ParamValidator interface {
	IsValidAggregation(value string) (model.ValidAggregation, bool)
	IsValidGrouping(value string) (model.ValidGrouping, bool)
	IsValidRegion(value string) (model.ValidRegion, bool)
	IsValidGenre(value string) (model.ValidGenre, bool)
	IsValidTimeSpan(startYear int, endYear int) bool
}

// Service that validates query parameters.
type QueryParamValidator struct {
	validator ParamValidator
}

// Creates a new instance of QueryParamValidator.
func NewQueryParamValidator(validator ParamValidator) QueryParamValidator {
	return QueryParamValidator{
		validator: validator,
	}
}

// Validates the contents of the provided data query struct.
func (q QueryParamValidator) Validate(query dto.DataQuery) (dto.DataRequest, error) {
	dataRequest := dto.DataRequest{}

	// Validate aggregate param
	aggregate, isValid := q.validator.IsValidAggregation(query.Aggregate)
	if !isValid {
		return dataRequest, &customerror.QueryParamError{Message: "Invalid aggregation param"}
	} else {
		dataRequest.Aggregation = aggregate
	}

	// If aggregation is by year or all games, validate year span
	if aggregate == model.AggregationByYear || aggregate == model.AggregationAllGames {
		startYear, err := strconv.Atoi(query.StartYear)
		if err != nil {
			return dataRequest, &customerror.QueryParamError{Message: "Invalid start-year format"}
		}
		endYear, err := strconv.Atoi(query.EndYear)
		if err != nil {
			return dataRequest, &customerror.QueryParamError{Message: "Invalid end-year format"}
		}
		validRange := q.validator.IsValidTimeSpan(startYear, endYear)
		if !validRange {
			return dataRequest, &customerror.QueryParamError{Message: "Invalid year range"}
		} else {
			dataRequest.StartYear = startYear
			dataRequest.EndYear = endYear
		}
	}

	// Validate grouping selection if not selection is all games
	if aggregate != model.AggregationAllGames {
		grouping, isValid := q.validator.IsValidGrouping(query.GroupBy)
		if !isValid {
			return dataRequest, &customerror.QueryParamError{Message: "Invalid grouping param"}
		} else {
			dataRequest.Group = grouping
		}

		// If grouping is by genre, validate genre
		if grouping == model.GroupingByGenre && aggregate == model.AggregationByYear {
			if len(query.Filters) == 0 {
				return dataRequest, &customerror.QueryParamError{Message: "No genre params supplied"}
			}
			for _, genre := range query.Filters {
				val, isValid := q.validator.IsValidGenre(genre)
				if !isValid {
					return dataRequest, &customerror.QueryParamError{Message: "Invalid genre param"}
				} else {
					dataRequest.Genres = append(dataRequest.Genres, val)
				}
			}
		}

		// If grouping is by region, validate region
		if grouping == model.GroupingByRegion && aggregate == model.AggregationByYear {
			if len(query.Filters) == 0 {
				return dataRequest, &customerror.QueryParamError{Message: "No region params supplied"}
			}

			for _, region := range query.Filters {
				val, isValid := q.validator.IsValidRegion(region)
				if !isValid {
					return dataRequest, &customerror.QueryParamError{Message: "Invalid region param"}
				} else {
					dataRequest.Regions = append(dataRequest.Regions, val)
				}
			}
		}
	}

	return dataRequest, nil
}
