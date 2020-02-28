package v2

// WiFi struct
type WiFi struct {
	MacAddress   string            `json:"apMac" mapstructure:"apMac"`
	Floors       []string          `json:"apFloors" mapstructure:"apFloors"`
	Tags         []string          `json:"apTags" mapstructure:"apTags"`
	Observations []WiFiObservation `json:"observations"`
}
