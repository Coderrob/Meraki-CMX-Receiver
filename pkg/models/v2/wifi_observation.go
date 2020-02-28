package v2

// WiFiObservation struct
type WiFiObservation struct {
	MacAddress      string   `json:"clientMac" mapstructure:"clientMac"`
	IPv4            string   `json:"ipv4"`
	IPv6            string   `json:"ipv6"`
	Location        Location `json:"location"`
	LastSeenAt      string   `json:"seenTime" mapstructure:"seenTime"`
	LastSeenAtEpoch int      `json:"seenEpoch" mapstructure:"seenEpoch"`
	SSID            string   `json:"ssid"`
	OS              string   `json:"os"`
	Manufacturer    string   `json:"manufacturer"`
	RSSI            int      `json:"rssi"`
}
