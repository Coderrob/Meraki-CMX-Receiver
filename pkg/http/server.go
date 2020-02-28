package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Server struct
type Server struct {
	*http.Server
}

// NewServer creates a new server instance
func NewServer(router *mux.Router) *Server {
	return &Server{&http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}}
}

// Start begins listening on a known address and port
func (server *Server) Start() {
	log.Fatal(server.ListenAndServe())
}
