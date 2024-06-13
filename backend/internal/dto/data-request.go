package dto

import "1dv027/wt2/internal/model"

// Holds validated query params used for requesting data from database/typesense server.
type DataRequest struct {
	Aggregation model.ValidAggregation
	Group       model.ValidGrouping
	StartYear   int
	EndYear     int
	Regions     []model.ValidRegion
	Genres      []model.ValidGenre
	Search      string
}
