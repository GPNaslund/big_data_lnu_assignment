package dataaccess

import (
	"1dv027/wt2/internal/dto"
	"1dv027/wt2/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// Struct for accessing clickhouse database and typesense server.
type VideoGamesDataAccess struct {
	db                 clickhouse.Conn
	typesenseSearchUrl string
	typesenseApiKey    string
	httpClient         *http.Client
}

// Creates a new instance of VideoGamesDataAccess.
func NewVideoGamesDataAccess(db clickhouse.Conn, typesenseSearchUrl string, typesenseApiKey string, httpClient *http.Client) VideoGamesDataAccess {
	return VideoGamesDataAccess{
		db:                 db,
		typesenseSearchUrl: typesenseSearchUrl,
		typesenseApiKey:    typesenseApiKey,
		httpClient:         httpClient,
	}
}

// Gets the total sales aggregated by region.
func (v VideoGamesDataAccess) GetAggregateTotalByRegions(ctx context.Context) ([]dto.SalesData, error) {
	query := `
        SELECT 'North America' AS Category, SUM(NaSales) AS Sales FROM video_games
        UNION ALL
        SELECT 'Europe' AS Category, SUM(EuSales) AS Sales FROM video_games
        UNION ALL
        SELECT 'Japan' AS Category, SUM(JpSales) AS Sales FROM video_games
        UNION ALL
        SELECT 'Other' AS Category, SUM(OtherSales) AS Sales FROM video_games
    `

	var results []dto.SalesData
	if err := v.db.Select(ctx, &results, query); err != nil {
		return nil, err
	}

	return results, nil
}

// Gets the total sales aggregated by genres.
func (v VideoGamesDataAccess) GetAggregateTotalByGenres(ctx context.Context) ([]dto.SalesData, error) {
	query := `
        SELECT Genre AS Category, SUM(GlobalSales) AS Sales FROM video_games GROUP BY Genre
    `

	var results []dto.SalesData
	if err := v.db.Select(ctx, &results, query); err != nil {
		return nil, err
	}

	return results, nil
}

// Gets the total sales for each selected year, grouped by region.
func (v VideoGamesDataAccess) GetAggregateByYearByRegions(ctx context.Context, startYear int, endYear int, regions []model.ValidRegion) ([]dto.YearlySalesData, error) {
	var queries []string
	var params []interface{}

	for _, region := range regions {
		regionColumn := v.getRegionCol(region)
		regionName := string(region)
		queryPart := fmt.Sprintf("SELECT Year, '%s' AS Category, SUM(%s) AS Sales FROM video_games WHERE Year BETWEEN ? AND ? GROUP BY Year", regionName, regionColumn)
		queries = append(queries, queryPart)
		params = append(params, startYear, endYear)
	}

	fullQuery := strings.Join(queries, " UNION ALL ") + " ORDER BY Year"

	var results []struct {
		Year     *int32   `db:"Year"`
		Category string   `db:"Category"`
		Sales    *float64 `db:"Sales"`
	}

	if err := v.db.Select(ctx, &results, fullQuery, params...); err != nil {
		return nil, err
	}

	yearlyDataMap := make(map[int32][]dto.SalesData)
	for _, result := range results {
		if result.Year != nil {
			yearlyDataMap[*result.Year] = append(yearlyDataMap[*result.Year], dto.SalesData{
				Category: result.Category,
				Sales:    *result.Sales,
			})
		}
	}

	var finalResults []dto.YearlySalesData
	for year, data := range yearlyDataMap {
		finalResults = append(finalResults, dto.YearlySalesData{
			Year: &year,
			Data: data,
		})
	}

	sort.Slice(finalResults, func(i, j int) bool {
		return *finalResults[i].Year < *finalResults[j].Year
	})

	return finalResults, nil
}

// Gets the total sales for each selected year, grouped by genre.
func (v VideoGamesDataAccess) GetAggregateByYearByGenres(ctx context.Context, startYear int, endYear int, genres []model.ValidGenre) ([]dto.YearlySalesData, error) {
	var results []struct {
		Year  *int32  `db:"Year"`
		Genre string  `db:"Genre"`
		Sales float64 `db:"TotalSales"`
	}

	genreStrings := make([]string, len(genres))
	for i, genre := range genres {
		genreStrings[i] = string(genre)
	}
	genreList := "'" + strings.Join(genreStrings, "','") + "'"

	query := fmt.Sprintf(`
        SELECT Year, Genre, SUM(GlobalSales) AS Sales
        FROM video_games
        WHERE Year BETWEEN ? AND ? AND Genre IN (%s)
        GROUP BY Year, Genre
        ORDER BY Year, Genre
    `, genreList)

	if err := v.db.Select(ctx, &results, query, startYear, endYear); err != nil {
		return nil, err
	}

	yearlyDataMap := make(map[int32][]dto.SalesData)
	for _, result := range results {
		if result.Year != nil {
			yearlyDataMap[*result.Year] = append(yearlyDataMap[*result.Year], dto.SalesData{
				Category: result.Genre,
				Sales:    result.Sales,
			})
		}
	}

	var finalResults []dto.YearlySalesData
	for year, data := range yearlyDataMap {
		finalResults = append(finalResults, dto.YearlySalesData{
			Year: &year,
			Data: data,
		})
	}

	sort.Slice(finalResults, func(i, j int) bool {
		return *finalResults[i].Year < *finalResults[j].Year
	})

	return finalResults, nil
}

// Gets all years released within selected time span.
func (v VideoGamesDataAccess) GetAggregateAllGames(ctx context.Context, startYear int, endYear int) ([]dto.VideoGame, error) {
	var games []dto.VideoGame
	query := `
        SELECT Rank, Name, Platform, Year, Genre, Publisher, 
               NaSales, EuSales, JpSales, OtherSales, GlobalSales
        FROM video_games
        WHERE Year BETWEEN ? AND ?
        ORDER BY Year, Rank
    `
	if err := v.db.Select(ctx, &games, query, startYear, endYear); err != nil {
		return nil, err
	}

	return games, nil
}

// Returnes data from typesense server based on search value.
func (v VideoGamesDataAccess) GetSearchData(ctx context.Context, search string, page int) (dto.TypesenseSearchResult, error) {
	emptyDto := dto.TypesenseSearchResult{}

	req, err := http.NewRequest("GET", v.typesenseSearchUrl, nil)
	if err != nil {
		return emptyDto, err
	}
	q := req.URL.Query()
	q.Add("q", search)
	q.Add("query_by", "Name,Platform,Genre")
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TYPESENSE-API-KEY", v.typesenseApiKey)

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return emptyDto, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return emptyDto, fmt.Errorf("Typesense server returned status: %d", resp.StatusCode)
	}

	var results dto.TypesenseSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return emptyDto, err
	}

	return results, nil
}

// Gets the database table column name based on provided region.
func (v VideoGamesDataAccess) getRegionCol(region model.ValidRegion) string {
	switch region {
	case model.RegionNorthAmerica:
		return "NaSales"
	case model.RegionEurope:
		return "EuSales"
	case model.RegionJapan:
		return "JpSales"
	case model.RegionOther:
		return "OtherSales"
	case model.RegionGlobal:
		return "GlobalSales"
	default:
		return "Unkown"
	}
}
