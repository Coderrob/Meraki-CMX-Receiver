package controllers

import (
	"coderrob.com/meraki-cmx-receiver/pkg/models"
	"coderrob.com/meraki-cmx-receiver/pkg/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type FeedController struct {
	OrganizationService services.IOrganizationService
	LocationService     services.ILocationService
}

func NewFeedController(
	organizationService services.IOrganizationService,
	locationService services.ILocationService) *FeedController {
	return &FeedController{
		OrganizationService: organizationService,
		LocationService:     locationService,
	}
}

func (controller *FeedController) AddRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/{orgID}", controller.ValidatorHandler).Methods("GET")
	router.HandleFunc("/{orgID}", controller.ReceiverHandler).Methods("POST")
	return router
}

func (controller *FeedController) ValidatorHandler(w http.ResponseWriter, r *http.Request) {
	organizationID := mux.Vars(r)["orgID"]
	validator, err := controller.OrganizationService.GetValidator(organizationID)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(validator))
}

func (controller *FeedController) ReceiverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil || r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()
	feed := &models.Feed{}
	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	organizationID := mux.Vars(r)["orgID"]
	secret, err := controller.OrganizationService.GetFeedSecret(organizationID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if secret != feed.Secret {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = controller.LocationService.UpdateLocations(feed)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
