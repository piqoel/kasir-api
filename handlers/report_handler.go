package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// multiple item apa aja, quantity nya
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	// Ambil query params (bisa kosong)
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	summary, err := h.service.GetReport(startDate, endDate)
	if err != nil {
		// Log error aslinya ke console server
		// fmt.Println("Report Error:", err)
		http.Error(w, "Gagal mengambil data report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
