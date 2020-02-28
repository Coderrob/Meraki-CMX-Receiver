package v3

// WiFi struct
type WiFi struct {
	NetworkId    string            `json:"networkId"`
	Observations []WiFiObservation `json:"observations"`
}
