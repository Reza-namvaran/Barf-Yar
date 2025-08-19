package handlers

import (
	"encoding/json"
	"encoding/csv"
	"time"
	"net/http"
	"strconv"
	_ "strings"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/models"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
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

// update
func (h *Handlers) UpdateActivityHandler(w http.ResponseWriter, r *http.Request) {
	var activity models.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.activityService.UpdateActivity(&activity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "activity updated successfully"})
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

func (h *Handlers) GetSupportersByActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activityID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	// Get activity details
	activity, err := h.activityService.GetActivityByID(activityID)
	if err != nil {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	}

	// Get supporters for this activity
	supporters, err := h.supporterService.GetSupportersByActivity(activityID)
	if err != nil {
		http.Error(w, "Failed to get supporters", http.StatusInternalServerError)
		return
	}

	data := templates.TemplateData{
		Title: "Activity Supporters",
		Data: map[string]interface{}{
			"Activity":   activity,
			"Supporters": supporters,
		},
	}

	w.Header().Set("Content-Type", "text/html")
	if err := h.templateService.RenderTemplate(w, "activity_supporters.html", data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) ExportSupporters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activityID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	format := vars["format"]

	supporters, err := h.supporterService.GetSupportersByActivity(activityID)
	if err != nil {
		http.Error(w, "Failed to get supporters", http.StatusInternalServerError)
		return
	}

	//TODO: Generate a filename with timestamp
	// timestamp := time.Now().Format("20060102_150405")
	// filename := fmt.Sprintf("supporters_activity_%d_%s.%s", activityID, timestamp, format)
	filename := "export"
	

	switch format {
	case "json":
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		h.exportJSON(w, supporters)
	case "csv":
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		h.exportCSV(w, supporters)
	case "pdf":
		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		h.exportPDF(w, supporters)
	default:
		http.Error(w, "Unsupported format", http.StatusBadRequest)
	}
}

func (h *Handlers) exportJSON(w http.ResponseWriter, supporters []*models.Supporter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(supporters)
}

func (h *Handlers) exportCSV(w http.ResponseWriter, supporters []*models.Supporter) {
	w.Header().Set("Content-Type", "text/csv")
	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"ID", "Activity ID", "User ID", "Joined At"})

	for _, s := range supporters {
		writer.Write([]string{
			strconv.Itoa(s.ID),
			strconv.Itoa(s.ActivityID),
			strconv.FormatInt(s.UserID, 10),
			s.JoinedAt.Format(time.RFC3339),
		})
	}
}

func (h *Handlers) exportPDF(w http.ResponseWriter, supporters []*models.Supporter) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Supporters Export")
	pdf.Ln(20)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(20, 10, "ID")
	pdf.Cell(30, 10, "Activity ID")
	pdf.Cell(40, 10, "User ID")
	pdf.Cell(50, 10, "Joined At")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	for _, s := range supporters {
		pdf.Cell(20, 10, strconv.Itoa(s.ID))
		pdf.Cell(30, 10, strconv.Itoa(s.ActivityID))
		pdf.Cell(40, 10, strconv.FormatInt(s.UserID, 10))
		pdf.Cell(50, 10, s.JoinedAt.Format("2006-01-02 15:04:05"))
		pdf.Ln(10)
	}

	w.Header().Set("Content-Type", "application/pdf")
	pdf.Output(w)
}