package model

// Represents a video game entity stored in the database.
type VideoGame struct {
	Rank        int32   `json:"Rank"`
	Name        string  `json:"Name"`
	Platform    string  `json:"Platform"`
	Year        int32   `json:"Year"`
	Genre       string  `json:"Genre"`
	Publisher   string  `json:"Publisher"`
	NaSales     float32 `json:"NaSales"`
	EuSales     float32 `json:"EuSales"`
	JpSales     float32 `json:"JpSales"`
	OtherSales  float32 `json:"OtherSales"`
	GlobalSales float32 `json:"GlobalSales"`
}
