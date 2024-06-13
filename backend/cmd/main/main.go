package main

import (
	"1dv027/wt2/internal/config"
	"1dv027/wt2/internal/database"
	"1dv027/wt2/internal/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Loads .env file from root
	envPath, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalf("Error getting absolute path to env")
	}
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Clickhouse database options
	options := clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", os.Getenv("CLICKHOUSE_HOST"), os.Getenv("CLICKHOUSE_PORT"))},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USERNAME"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
	}

	db := database.NewClickHouseDatabase()
	conn, err := db.NewDatabaseConnection(options)
	if err != nil {
		log.Fatalf("Error creating database connection: %+v", err)
	}

	client := &http.Client{}
	// Config for IoC-Container
	containerConfig := config.ContainerConfig{
		DbConnection:       conn,
		ApiKey:             os.Getenv("API_KEY"),
		TypesenseSearchUrl: os.Getenv("TYPESENSE_SEARCH_URL"),
		TypesenseApiKey:    os.Getenv("TYPESENSE_API_KEY"),
		HttpClient:         client,
	}
	container := config.SetupContainer(containerConfig)

	router := router.NewRouter(container)

	router.StartRouter()
}
