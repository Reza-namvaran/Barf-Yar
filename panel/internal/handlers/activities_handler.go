package handlers

import (
	"net/http"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
)

func (h *Handlers) GetAllActivities(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO Replace with middleware
	cookie, err := r.Cookie("session_token")
	if err != nil || !h.authService.ValidateSessionToken(cookie.Value) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	//==================================================================

	activities, err := h.activityService.GetAllActivities()
	if err != nil {
		http.Error(w, "Failed to fetch activities", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	data := templates.TemplateData{
		Title: "Activities",
		Data:  activities,
	}

	if err := h.templateService.RenderTemplate(w, "activities.html", data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
