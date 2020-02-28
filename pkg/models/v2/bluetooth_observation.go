package v2

// BluetoothObservation struct
type BluetoothObservation struct {
	MacAddress      string   `json:"clientMac" mapstructure:"clientMac"`
	Location        Location `json:"location"`
	LastSeenAt      string   `json:"seenTime" mapstructure:"seenTime"`
	LastSeenAtEpoch int      `json:"seenEpoch" mapstructure:"seenEpoch"`
	RSSI            int      `json:"rssi"`
}
