package model

import "encoding/json"

// Generic entity for raw json data.
type Dataset struct {
	Data json.RawMessage `json:"data"`
}
