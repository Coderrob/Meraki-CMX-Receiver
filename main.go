package main

import (
	"coderrob.com/meraki-cmx-receiver/pkg/services"
	"github.com/gorilla/mux"
)

func main() {
	config := map[string]services.Settings{
		"78669f65-07d2-40a7-9d6c-1974934c508d": {
			Secret:    "supersecret",
			Validator: "i_am_totally_valid",
		},
	}
	feedController, _ := InitializeFeedController(config)
	router := mux.NewRouter()
	feedController.AddRoutes(router)
	server, _ := InitializeServer(router)
	server.Start()
}
