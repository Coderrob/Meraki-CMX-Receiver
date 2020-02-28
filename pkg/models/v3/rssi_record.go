package v3

// RSSIRecord struct
type RSSIRecord struct {
	ApMacAddress string `json:"apMac" mapstructure:"ApMacAddress"`
	RSSI         int    `json:"rssi"`
}
