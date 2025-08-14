package handlers

import (
	"encoding/json"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/storage"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
	"net/http"
	"strconv"
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

// add
func (h *Handlers) AddActivityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var activity storage.Activity

	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.activityService.AddActivity(&activity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "activity added successfully"})

}

// delete
func (h *Handlers) DeleteActivityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Path[len("/dashboard/activities/delete"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}
	if err := h.activityService.DeleteActivity(id); err != nil {
		http.Error(w, "faild to delete activity", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "activity deleted successfully"})
}
