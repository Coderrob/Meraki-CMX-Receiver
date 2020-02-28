package v2

// Bluetooth struct
type Bluetooth struct {
	MacAddress   string                 `json:"apMac" mapstructure:"apMac"`
	Floors       []string               `json:"apFloors" mapstructure:"apFloors"`
	Tags         []string               `json:"apTags" mapstructure:"apTags"`
	Observations []BluetoothObservation `json:"observations"`
}
