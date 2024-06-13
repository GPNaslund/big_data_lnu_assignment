package videogamesdata

import (
	"1dv027/wt2/internal/model"
	"context"
	"log"
	"path/filepath"
	"strconv"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/xuri/excelize/v2"
)

// Struct for creating database table and seed it with data from dataset.
type VideoGamesSeeder struct {
	db clickhouse.Conn
}

// Creates a new video games seeder.
func NewVideoGamesSeeder(db clickhouse.Conn) VideoGamesSeeder {
	return VideoGamesSeeder{
		db: db,
	}
}

// Method for creating a table.
func (v VideoGamesSeeder) CreateTable(ctx context.Context) error {
	query := `
    CREATE TABLE IF NOT EXISTS video_games_db.video_games (
        Rank Int32,
        Name String,
        Platform String,
        Year Int32,
        Genre String,
        Publisher String,
        NaSales Float32,
        EuSales Float32,
        JpSales Float32,
        OtherSales Float32,
        GlobalSales Float32
    )
    ENGINE = MergeTree
    ORDER BY (Year, Platform, Genre, Publisher)
  `

	err := v.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

// Method for seeding database table with data from dataset.
func (v VideoGamesSeeder) SeedVideoGamesSalesData(ctx context.Context) error {
	videoGames := v.readVideoGamesSalesDataFromFile()
	batch, err := v.db.PrepareBatch(ctx, "INSERT INTO video_games")
	if err != nil {
		return err
	}
	for _, videoGame := range videoGames {
		err := batch.Append(
			videoGame.Rank,
			videoGame.Name,
			videoGame.Platform,
			videoGame.Year,
			videoGame.Genre,
			videoGame.Publisher,
			videoGame.NaSales,
			videoGame.EuSales,
			videoGame.JpSales,
			videoGame.OtherSales,
			videoGame.GlobalSales,
		)
		if err != nil {
			return err
		}
	}
	return batch.Send()
}

// Reads the dataset and constructs a slice of model entities from the content.
func (v VideoGamesSeeder) readVideoGamesSalesDataFromFile() []model.VideoGame {
	absolutePath, err := filepath.Abs("./db-seed/data/video-games/video_games_sales.xlsx")
	if err != nil {
		log.Fatalf("failed to get absolute path")
	}

	f, err := excelize.OpenFile(absolutePath)
	if err != nil {
		log.Fatal("failed to open excel file:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("failed to close file:", err)
		}
	}()

	var videoGamesSales []model.VideoGame
	rows, err := f.GetRows("video_games_sales")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		rank, err := strconv.Atoi(row[0])
		if err != nil {
			log.Print("failed to convert rank to int")
			rank = 0
		}
		name := row[1]
		platform := row[2]
		year, err := strconv.Atoi(row[3])
		if err != nil {
			log.Print("failed to convert year to int")
			year = 0
		}
		genre := row[4]
		publisher := row[5]
		naSales, err := strconv.ParseFloat(row[6], 32)
		if err != nil {
			log.Fatal("failed to convert na-sales to float")
		}
		euSales, err := strconv.ParseFloat(row[7], 32)
		if err != nil {
			log.Fatal("failed to convert eu-sales to float")
		}
		jpSales, err := strconv.ParseFloat(row[8], 32)
		if err != nil {
			log.Fatal("failed to convert jp-sales to float")
		}
		otherSales, err := strconv.ParseFloat(row[9], 32)
		if err != nil {
			log.Fatal("failed to convert other-sales to float")
		}
		globalSales, err := strconv.ParseFloat(row[10], 32)
		if err != nil {
			log.Fatal("failed to convert global-sales to float")
		}

		videoGame := model.VideoGame{
			Rank:        int32(rank),
			Name:        name,
			Platform:    platform,
			Year:        int32(year),
			Genre:       genre,
			Publisher:   publisher,
			NaSales:     float32(naSales),
			EuSales:     float32(euSales),
			JpSales:     float32(jpSales),
			OtherSales:  float32(otherSales),
			GlobalSales: float32(globalSales),
		}
		videoGamesSales = append(videoGamesSales, videoGame)
	}
	return videoGamesSales
}
