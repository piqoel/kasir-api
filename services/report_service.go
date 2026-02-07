package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport(startDateStr, endDateStr string) (models.SalesSummary, error) {
	var startTime, endTime time.Time

	// Logic: Jika parameter kosong, asumsikan "Hari Ini"
	if startDateStr == "" || endDateStr == "" {
		now := time.Now()
		// Start: Jam 00:00:00 hari ini
		startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		// End: Jam 23:59:59 hari ini
		endTime = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	} else {
		// Challenge Logic: Parse tanggal dari query param (Format YYYY-MM-DD)
		var err error
		layout := "2006-01-02" // Magic date Go untuk format YYYY-MM-DD

		startTime, err = time.Parse(layout, startDateStr)
		if err != nil {
			return models.SalesSummary{}, err
		}

		parsedEnd, err := time.Parse(layout, endDateStr)
		if err != nil {
			return models.SalesSummary{}, err
		}
		// Set end date ke akhir hari tersebut (23:59:59)
		endTime = time.Date(parsedEnd.Year(), parsedEnd.Month(), parsedEnd.Day(), 23, 59, 59, 0, parsedEnd.Location())
	}

	return s.repo.GetSalesSummary(startTime, endTime)
}
