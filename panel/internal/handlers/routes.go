package handlers

import (
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/middleware"
    "github.com/gorilla/mux"
	"net/http"
)

// SetupRoutes configures all HTTP routes for the application
func SetupRoutes(handlers *Handlers) *mux.Router {

	new_router := mux.NewRouter()

	// Static file serving
	fs := http.FileServer(http.Dir("static"))
	new_router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// API routes
	api := new_router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", handlers.Login).Methods("Post")
	api.HandleFunc("/logout", handlers.Logout).Methods("Post")
	
	// Dashboard routs
	dashBoard := new_router.PathPrefix("/dashboard").Subrouter()
	dashBoard.Use(middleware.AuthMiddleware(handlers.authService))
	dashBoard.HandleFunc("", handlers.Dashboard).Methods("Get")
	
	activities := dashBoard.PathPrefix("/activities").Subrouter()
	activities.HandleFunc("", handlers.GetAllActivities).Methods("Get")
	activities.HandleFunc("/", handlers.AddActivityHandler).Methods("Post")
	activities.HandleFunc("/{id}", handlers.DeleteActivityHandler).Methods("Delete")
	activities.HandleFunc("/{id}", handlers.UpdateActivityHandler).Methods("Put")
	activities.HandleFunc("/{id}/supporters", handlers.GetSupportersByActivity).Methods("GET")
	activities.HandleFunc("/{id}/supporters/export/{format}", handlers.ExportSupporters).Methods("GET")

	// Page routes
	new_router.PathPrefix("/").HandlerFunc(handlers.LoginPage).Methods("Get")

	return new_router
}
