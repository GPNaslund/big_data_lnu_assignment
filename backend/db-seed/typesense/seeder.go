package typesense

import (
	"1dv027/wt2/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type CollectionSchema struct {
	Name   string `json:"name"`
	Fields []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"fields"`
}

// Struct for creating a schema and seeding the typesense server.
type TypesenseDataSeeder struct {
	apiKey string
	host   string
}

// Creates an instance of TypesenseDataSeeder.
func NewTypesenseDataSeeder(apiKey string, host string) TypesenseDataSeeder {
	return TypesenseDataSeeder{
		apiKey: apiKey,
		host:   host,
	}
}

// Method for creating a collection.
func (t TypesenseDataSeeder) CreateCollection() error {
	schema := CollectionSchema{
		Name: "video_games_sales",
		Fields: []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		}{
			{Name: "Rank", Type: "int32"},
			{Name: "Name", Type: "string"},
			{Name: "Platform", Type: "string"},
			{Name: "Year", Type: "int32"},
			{Name: "Genre", Type: "string"},
			{Name: "Publisher", Type: "string"},
			{Name: "NaSales", Type: "float"},
			{Name: "EuSales", Type: "float"},
			{Name: "JpSales", Type: "float"},
			{Name: "OtherSales", Type: "float"},
			{Name: "GlobalSales", Type: "float"},
		},
	}

	data, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/collections", t.host)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-TYPESENSE-API-KEY", t.apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println("Collection created:", result)

	return nil
}

// Seeds typesense server from clickhouse database.
func (t TypesenseDataSeeder) SeedClickhouseData(ctx context.Context, db clickhouse.Conn) error {
	rows, err := db.Query(ctx, "SELECT * FROM video_games")
	if err != nil {
		log.Printf("Error querying ClickHouse: %v", err)
		return err
	}
	defer rows.Close()

	client := &http.Client{}

	for rows.Next() {
		var game model.VideoGame
		if err := rows.Scan(&game.Rank, &game.Name, &game.Platform, &game.Year, &game.Genre, &game.Publisher, &game.NaSales, &game.EuSales, &game.JpSales, &game.OtherSales, &game.GlobalSales); err != nil {
			log.Printf("Error scanning game data: %v", err)
			return err
		}

		jsonData, err := json.Marshal(game)
		if err != nil {
			log.Printf("Error marshaling game data to JSON: %v", err)
			return err
		}

		request, err := http.NewRequest("POST", fmt.Sprintf("%s/collections/%s/documents", t.host, "video_games_sales"), strings.NewReader(string(jsonData)))
		if err != nil {
			log.Printf("Error creating request: %v", err)
			return err
		}
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-TYPESENSE-API-KEY", t.apiKey)

		response, err := client.Do(request)
		if err != nil {
			log.Printf("Error sending data to Typesense: %v", err)
			return err
		}

		response.Body.Close()

	}

	if err := rows.Err(); err != nil {
		log.Printf("Error after iterating rows: %v", err)
		return err
	}

	log.Println("Successfully seeded Typesense")
	return nil
}
