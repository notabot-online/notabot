package router

import (
	"github.com/gorilla/mux"
	"github.com/notabot/backend/internal/api/handler"
	"github.com/notabot/backend/internal/api/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)
	r.Use(middleware.Recovery)

	// Health check
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.Auth) // Apply auth middleware to API routes

	// Verification routes
	verification := api.PathPrefix("/verification").Subrouter()
	verification.HandleFunc("", handler.CreateVerification).Methods("POST")
	verification.HandleFunc("/{id}", handler.GetVerification).Methods("GET")
	verification.HandleFunc("/{id}/proof", handler.SubmitProof).Methods("POST")

	return r
}


