package model

type ValidAggregation int

const (
	AggregationByYear ValidAggregation = iota
	AggregationTotal
	AggregationAllGames
)

// Defines valid aggregations.
var validAggregations = map[string]ValidAggregation{
	"by year":   AggregationByYear,
	"total":     AggregationTotal,
	"all games": AggregationAllGames,
}

type ValidGrouping int

const (
	GroupingByRegion ValidGrouping = iota
	GroupingByGenre
)

// Defines valid groupings.
var validGroupings = map[string]ValidGrouping{
	"by region": GroupingByRegion,
	"by genre":  GroupingByGenre,
}

type ValidRegion string

const (
	RegionNorthAmerica ValidRegion = "North America"
	RegionEurope       ValidRegion = "Europe"
	RegionJapan        ValidRegion = "Japan"
	RegionOther        ValidRegion = "Other"
	RegionGlobal       ValidRegion = "Global"
)

// Defines valid regions
var validRegions = map[string]ValidRegion{
	"north america": RegionNorthAmerica,
	"europe":        RegionEurope,
	"japan":         RegionJapan,
	"other":         RegionOther,
	"global":        RegionGlobal,
}

type ValidGenre string

const (
	GenreAction      ValidGenre = "Action"
	GenrePuzzle      ValidGenre = "Puzzle"
	GenreSports      ValidGenre = "Sports"
	GenreAdventure   ValidGenre = "Adventure"
	GenreRacing      ValidGenre = "Racing"
	GenreRolePlaying ValidGenre = "Role-Playing"
	GenreShooter     ValidGenre = "Shooter"
	GenreSimulation  ValidGenre = "Simulation"
	GenreMisc        ValidGenre = "Misc"
	GenrePlatform    ValidGenre = "Platform"
	GenreFighting    ValidGenre = "Fighting"
	GenreStrategy    ValidGenre = "Strategy"
)

// Defines valid genres
var validGenres = map[string]ValidGenre{
	"action":       GenreAction,
	"puzzle":       GenrePuzzle,
	"sports":       GenreSports,
	"adventure":    GenreAdventure,
	"racing":       GenreRacing,
	"role-playing": GenreRolePlaying,
	"shooter":      GenreShooter,
	"simulation":   GenreSimulation,
	"misc":         GenreMisc,
	"platform":     GenrePlatform,
	"fighting":     GenreFighting,
	"strategy":     GenreStrategy,
}

// Defines valid year range
var validYearRange = map[string]int{
	"start-year": 1977,
	"end-year":   2020,
}

// Struct that holds all valid query parameters, and methods
// for validating input against valid parameters.
type DataParameters struct {
	aggregations map[string]ValidAggregation
	groupings    map[string]ValidGrouping
	regions      map[string]ValidRegion
	genres       map[string]ValidGenre
	yearRange    map[string]int
}

// Creates a new instance of DataParameters entity.
func NewDataParameters() DataParameters {
	return DataParameters{
		aggregations: validAggregations,
		groupings:    validGroupings,
		regions:      validRegions,
		genres:       validGenres,
		yearRange:    validYearRange,
	}
}

// Getter for all valid aggregations.
func (d DataParameters) GetAggregations() []string {
	var result []string
	for key := range d.aggregations {
		result = append(result, key)
	}
	return result
}

// Getter for all valid types of groupings.
func (d DataParameters) GetGroupings() []string {
	var result []string
	for key := range d.groupings {
		result = append(result, key)
	}
	return result
}

// Getter for all valid regions.
func (d DataParameters) GetRegions() []string {
	var result []string
	for key := range d.regions {
		result = append(result, key)
	}
	return result
}

// Getter for all valid genres.
func (d DataParameters) GetGenres() []string {
	var result []string
	for key := range d.genres {
		result = append(result, key)
	}
	return result
}

// Getter for valid start year.
func (d DataParameters) GetStartYear() int {
	return d.yearRange["start-year"]
}

// Getter for valid end year.
func (d DataParameters) GetEndYear() int {
	return d.yearRange["end-year"]
}

// Validates if input string is a valid aggregation.
func (d DataParameters) IsValidAggregation(value string) (ValidAggregation, bool) {
	val, exists := d.aggregations[value]
	return val, exists
}

// Validates if input string is a valid grouping.
func (d DataParameters) IsValidGrouping(value string) (ValidGrouping, bool) {
	val, exists := d.groupings[value]
	return val, exists
}

// Validates if input string is a valid region.
func (d DataParameters) IsValidRegion(value string) (ValidRegion, bool) {
	val, exists := d.regions[value]
	return val, exists
}

// Validates if input string is a valid genre.
func (d DataParameters) IsValidGenre(value string) (ValidGenre, bool) {
	val, exists := d.genres[value]
	return val, exists
}

// Validates if the provided startYear and endYear is within the valid year span.
func (d DataParameters) IsValidTimeSpan(startYear int, endYear int) bool {
	return startYear >= d.yearRange["start-year"] && endYear <= d.yearRange["end-year"]
}
