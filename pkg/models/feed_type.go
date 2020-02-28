package models

// Enumeration values for Meraki location feed types
const (
	WiFi                 FeedType = "WiFi"
	WiFiDevicesSeen      FeedType = "DevicesSeen"
	BluetoothDevicesSeen FeedType = "BluetoothDevicesSeen"
)

// FeedType struct
type FeedType string

// String returns the feed type description
func (ft FeedType) String() string {
	return string(ft)
}
