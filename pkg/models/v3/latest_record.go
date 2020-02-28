package v3

// LatestRecord struct
type LatestRecord struct {
	NearestApMacAddress string `json:"nearestApMac" mapstructure:"nearestApMac"`
	NearestApRSSI       string `json:"nearestApRssi" mapstructure:"nearestApRssi"`
	Time                string `json:"time"`
}
