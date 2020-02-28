package v2

// Location struct
type Location struct {
	Latitude    float64   `json:"lat" mapstructure:"lat"`
	Longitude   float64   `json:"lng" mapstructure:"lng"`
	Uncertainty float64   `json:"unc" mapstructure:"unc"`
	X           []float64 `json:"x"`
	Y           []float64 `json:"y"`
}
