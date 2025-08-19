package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	_ "strings"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
	"github.com/gorilla/mux"
)

func (h *Handlers) GetAllActivities(w http.ResponseWriter, r *http.Request) {
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
	var activity models.Activity

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
	vars := mux.Vars(r)
	idStr := vars["id"]

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