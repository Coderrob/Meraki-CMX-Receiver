//+build wireinject

package main

import (
	"coderrob.com/meraki-cmx-receiver/pkg/http"
	"coderrob.com/meraki-cmx-receiver/pkg/http/controllers"
	"coderrob.com/meraki-cmx-receiver/pkg/services"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

// ServiceSet is the set of common services used by controllers
var ServiceSet = wire.NewSet(
	services.NewOrganizationService,
	services.NewLocationService)

// InitializeFeedController initializes and returns new instance of FeedController
func InitializeFeedController(config map[string]services.Settings) (*controllers.FeedController, error) {
	wire.Build(
		ServiceSet,
		controllers.NewFeedController)
	return &controllers.FeedController{}, nil
}

// InitializeServer initializes and returns new instance of http.Server
func InitializeServer(router *mux.Router) (*http.Server, error) {
	wire.Build(http.NewServer)
	return &http.Server{}, nil
}
