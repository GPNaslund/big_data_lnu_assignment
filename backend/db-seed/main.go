package main

import (
	videogamesdata "1dv027/wt2/db-seed/data/video-games"
	"1dv027/wt2/db-seed/typesense"
	"1dv027/wt2/internal/database"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/joho/godotenv"
)

// Seeds database and typesense server with data.
func main() {

	// Get path to .env and loads it.
	envPath, err := filepath.Abs(".env")
	if err != nil {
		log.Fatalf("Error getting absolute path to env")
	}
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Sets up clickhouse database config.
	options := clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", os.Getenv("CLICKHOUSE_HOST"), os.Getenv("CLICKHOUSE_PORT"))},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USERNAME"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
	}

	// Creates a new connection to the clickhouse database.
	db := database.NewClickHouseDatabase()
	conn, err := db.NewDatabaseConnection(options)
	if err != nil {
		log.Fatalf("Error creating database connection: %+v", err)
	}

	// Creates a new video games seeder for setting up the database.
	dbSeeder := videogamesdata.NewVideoGamesSeeder(conn)
	ctx := context.Background()

	// Creates database table.
	err = dbSeeder.CreateTable(ctx)
	if err != nil {
		log.Fatalf("Error creating table: %+v", err)
	}

	// Seeds database with data.
	err = dbSeeder.SeedVideoGamesSalesData(ctx)
	if err != nil {
		log.Fatalf("Error seeding video games data: %+v", err)
	}

	// Sets up typesense database seeder.
	typeSenseApiKey := os.Getenv("TYPESENSE_API_KEY")
	typesenseHost := os.Getenv("TYPESENSE_HOST")
	typesenseSeeder := typesense.NewTypesenseDataSeeder(typeSenseApiKey, typesenseHost)

	// Creates the typesense collection.
	err = typesenseSeeder.CreateCollection()
	if err != nil {
		log.Fatalf("Error creating collection: %s", err.Error())
	}

	// Seeds typesense server with data from clickhouse database.
	err = typesenseSeeder.SeedClickhouseData(ctx, conn)
	if err != nil {
		log.Fatalf("Error seeding typsense server with clickhouse data: %s", err)
	}
}
