package v3

// Location struct
type Location struct {
	FloorPlanID   string       `json:"floorPlanId"`
	FloorPlanName string       `json:"floorPlanName"`
	Time          string       `json:"time"`
	Latitude      float64      `json:"lat" mapstructure:"lat"`
	Longitude     float64      `json:"lng" mapstructure:"lng"`
	Variance      float64      `json:"variance"`
	Tags          []string     `json:"nearestApTags" mapstructure:"nearestApTags"`
	RSSIRecords   []RSSIRecord `json:"rssiRecords"`
	X             float64      `json:"x"`
	Y             float64      `json:"y"`
}
