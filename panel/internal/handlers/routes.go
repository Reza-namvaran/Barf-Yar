package handlers

import (
	"net/http"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/middleware"
)

// SetupRoutes configures all HTTP routes for the application
func SetupRoutes(handlers *Handlers) {
	// Static file serving
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// API routes
	http.HandleFunc("/api/login", handlers.Login)
	http.HandleFunc("/api/logout", handlers.Logout)

	// Page routes
	http.HandleFunc("/", handlers.LoginPage)
	http.Handle("/dashboard", 
		middleware.AuthMiddleware(handlers.authService)(
			http.HandlerFunc(handlers.Dashboard),
		),
	)
	http.HandleFunc("/activities", handlers.Activities)
}
