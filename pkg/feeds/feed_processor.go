package feeds

import (
	"coderrob.com/meraki-cmx-receiver/pkg/models"
	v2 "coderrob.com/meraki-cmx-receiver/pkg/models/v2"
	v3 "coderrob.com/meraki-cmx-receiver/pkg/models/v3"
	"errors"
	"github.com/mitchellh/mapstructure"
	"time"
)

type feedProcessor func(interface{}) ([]ClientLocation, error)
type feedProcessorMap map[models.FeedType]feedProcessor

// Processor the feed type mapped to a feed processor function
var Processor = feedProcessorMap{
	models.WiFiDevicesSeen:      processWiFiDevicesSeenType,
	models.WiFi:                 processWiFiType,
	models.BluetoothDevicesSeen: processBluetoothDevicesSeenType,
}

// Process takes the feed data and processes the json data to a known type
func Process(feed *models.Feed) ([]ClientLocation, error) {
	if processor, exists := Processor[feed.Type]; exists {
		return processor(feed.Data)
	}
	return nil, errors.New("feed type not supported")
}

func processWiFiType(data interface{}) ([]ClientLocation, error) {
	wifi := &v3.WiFi{}
	_ = mapstructure.Decode(data, wifi)
	var clientLocations = make([]ClientLocation, 0)
	for index := range wifi.Observations {
		mostRecentTimeSeen := time.Time{}
		currentLocationIndex := 0
		if wifi.Observations[index].Locations == nil {
			continue
		}
		for locIndex := range wifi.Observations[index].Locations {
			seenAt, _ := time.Parse(time.RFC3339, wifi.Observations[index].Locations[locIndex].Time)
			if mostRecentTimeSeen.Before(seenAt) {
				mostRecentTimeSeen = seenAt
				currentLocationIndex = locIndex
			}
		}
		clientLocations = append(clientLocations, ClientLocation{
			MacAddress:  wifi.Observations[index].MacAddress,
			LastSeenAt:  mostRecentTimeSeen,
			Latitude:    wifi.Observations[index].Locations[currentLocationIndex].Latitude,
			Longitude:   wifi.Observations[index].Locations[currentLocationIndex].Longitude,
			XCoordinate: wifi.Observations[index].Locations[currentLocationIndex].X,
			YCoordinate: wifi.Observations[index].Locations[currentLocationIndex].Y,
		})
	}
	return clientLocations, nil
}

func processWiFiDevicesSeenType(data interface{}) ([]ClientLocation, error) {
	wifi := &v2.WiFi{}
	_ = mapstructure.Decode(data, wifi)
	var clientLocations = make([]ClientLocation, 0)
	for index := range wifi.Observations {
		seenAt, _ := time.Parse(time.RFC3339, wifi.Observations[index].LastSeenAt)
		clientLocations = append(clientLocations, ClientLocation{
			MacAddress:  wifi.Observations[index].MacAddress,
			LastSeenAt:  seenAt,
			Latitude:    wifi.Observations[index].Location.Latitude,
			Longitude:   wifi.Observations[index].Location.Longitude,
			XCoordinate: 0,
			YCoordinate: 0,
		})
	}
	return clientLocations, nil
}

func processBluetoothDevicesSeenType(data interface{}) ([]ClientLocation, error) {
	bluetooth := &v2.Bluetooth{}
	_ = mapstructure.Decode(data, bluetooth)
	var clientLocations = make([]ClientLocation, 0)
	for index := range bluetooth.Observations {
		seenAt, _ := time.Parse(time.RFC3339, bluetooth.Observations[index].LastSeenAt)
		clientLocations = append(clientLocations, ClientLocation{
			MacAddress:  bluetooth.Observations[index].MacAddress,
			LastSeenAt:  seenAt,
			Latitude:    bluetooth.Observations[index].Location.Latitude,
			Longitude:   bluetooth.Observations[index].Location.Longitude,
			XCoordinate: 0,
			YCoordinate: 0,
		})
	}
	return clientLocations, nil
}
