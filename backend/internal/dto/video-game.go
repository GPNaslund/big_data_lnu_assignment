package dto

// Holds video game data.
type VideoGame struct {
	Rank        int32   `ch:"Rank"`
	Name        string  `ch:"Name"`
	Platform    string  `ch:"Platform"`
	Year        int32   `ch:"Year"`
	Genre       string  `ch:"Genre"`
	Publisher   string  `ch:"Publisher"`
	NaSales     float32 `ch:"NaSales"`
	EuSales     float32 `ch:"EuSales"`
	JpSales     float32 `ch:"JpSales"`
	OtherSales  float32 `ch:"OtherSales"`
	GlobalSales float32 `ch:"GlobalSales"`
}
