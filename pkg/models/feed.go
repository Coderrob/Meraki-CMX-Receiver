package models

// Feed struct
type Feed struct {
	Version Version     `json:"version"`
	Secret  string      `json:"secret"`
	Type    FeedType    `json:"type"`
	Data    interface{} `json:"data"`
}
