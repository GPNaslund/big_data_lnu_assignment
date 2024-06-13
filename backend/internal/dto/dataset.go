package dto

import "encoding/json"

// Generic struct for holding raw json.
type Dataset struct {
	Data json.RawMessage `json:"data"`
}
