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
	api.HandleFunc("/login", handlers.Login)
	api.HandleFunc("/logout", handlers.Logout)
	
	// Dashboard routs
	dashBoard := new_router.PathPrefix("/dashboard").Subrouter()
	dashBoard.Use(middleware.AuthMiddleware(handlers.authService))
	dashBoard.HandleFunc("", handlers.Dashboard)
	
	activities := dashBoard.PathPrefix("/activities").Subrouter()
	activities.HandleFunc("", handlers.GetAllActivities)
	activities.HandleFunc("/add", handlers.AddActivityHandler)
	activities.HandleFunc("/delete/{id}", handlers.DeleteActivityHandler)

	// Page routes
	new_router.PathPrefix("/").HandlerFunc(handlers.LoginPage)

	return new_router
}
