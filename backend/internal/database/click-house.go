package database

import "github.com/ClickHouse/clickhouse-go/v2"

// Struct for creating a connection to a clickhouse database.
type ClickHouseDatabase struct{}

// Creates a new instance of a clickhouse database.
func NewClickHouseDatabase() ClickHouseDatabase {
	return ClickHouseDatabase{}
}

// Creates and returns a connection based on provided options.
func (c ClickHouseDatabase) NewDatabaseConnection(options clickhouse.Options) (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&options)
	if err != nil {
		return nil, err
	}
	return conn, err
}
