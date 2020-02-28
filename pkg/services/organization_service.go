package services

import (
	"fmt"
)

// Settings struct
type Settings struct {
	Secret    string
	Validator string
}

// IOrganizationService interface
type IOrganizationService interface {
	GetValidator(organizationId string) (string, error)
	GetFeedSecret(organizationId string) (string, error)
}

// OrganizationService struct
type OrganizationService struct {
	IOrganizationService
	Config map[string]Settings
}

// NewOrganizationService returns new instance of IOrganizationService
func NewOrganizationService(config map[string]Settings) IOrganizationService {
	return &OrganizationService{Config: config}
}

// GetValidator returns the Meraki dashboard locations feed validator
func (svc OrganizationService) GetValidator(organizationId string) (string, error) {
	if settings, exists := svc.Config[organizationId]; exists {
		return settings.Validator, nil
	}
	return "", fmt.Errorf("settings not found for organization %s", organizationId)
}

// GetFeedSecret returns the Meraki dashboard locations feed secret for verifying feed payloads
func (svc OrganizationService) GetFeedSecret(organizationId string) (string, error) {
	if settings, exists := svc.Config[organizationId]; exists {
		return settings.Secret, nil
	}
	return "", fmt.Errorf("settings not found for organization %s", organizationId)
}
