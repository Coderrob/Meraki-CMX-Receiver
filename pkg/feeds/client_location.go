package feeds

import "time"

// ClientLocation struct
type ClientLocation struct {
	MacAddress  string
	LastSeenAt  time.Time
	Latitude    float64
	Longitude   float64
	XCoordinate float64
	YCoordinate float64
}
