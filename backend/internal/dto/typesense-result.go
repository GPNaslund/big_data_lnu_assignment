package dto

// Holds data response from typesense server query.
type TypesenseSearchResult struct {
	FacetCounts []interface{} `json:"facet_counts"`
	Found       int           `json:"found"`
	Hits        []struct {
		Highlights []struct {
			Field   string `json:"field"`
			Snippet string `json:"snippet"`
		} `json:"highlights"`
		Document struct {
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
		} `json:"document"`
	}
}
