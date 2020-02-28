package services

import (
	"coderrob.com/meraki-cmx-receiver/pkg/feeds"
	"coderrob.com/meraki-cmx-receiver/pkg/models"
)

// ILocationService interface
type ILocationService interface {
	UpdateLocations(feed *models.Feed) error
}

// LocationService struct
type LocationService struct {
	ILocationService
}

// NewLocationService returns new instance of ILocationService
func NewLocationService() ILocationService {
	return &LocationService{}
}

// UpdateLocations processes the locations feed to update device locations
func (service *LocationService) UpdateLocations(feed *models.Feed) error {
	locations, err := feeds.Process(feed)
	for range locations {
		// TODO - here is where the device locations would be used
	}
	return err
}
