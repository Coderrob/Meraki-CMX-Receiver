package v3

// WiFiObservation struct
type WiFiObservation struct {
	MacAddress   string       `json:"clientMac" mapstructure:"clientMac"`
	IPv4         string       `json:"ipv4"`
	IPv6         string       `json:"ipv6"`
	Locations    []Location   `json:"feeds"`
	SSID         string       `json:"ssid"`
	OS           string       `json:"os"`
	Manufacturer string       `json:"manufacturer"`
	LatestRecord LatestRecord `json:"latestRecord"`
}
