package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSalesSummary(startDate, endDate time.Time) (models.SalesSummary, error) {
	var summary models.SalesSummary

	// 1. Query Total Revenue & Total Transaksi
	queryStats := `
		SELECT 
			COALESCE(SUM(total_amount), 0), 
			COUNT(id) 
		FROM transactions 
		WHERE created_at >= $1 AND created_at <= $2`

	err := repo.db.QueryRow(queryStats, startDate, endDate).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)
	if err != nil {
		return summary, err
	}

	// 2. Query Produk Terlaris
	queryBestSeller := `
		SELECT 
			p.name, 
			SUM(td.quantity) as total_qty 
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name 
		ORDER BY total_qty DESC 
		LIMIT 1`

	err = repo.db.QueryRow(queryBestSeller, startDate, endDate).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	// Handle logic: Jika tidak ada data penjualan sama sekali
	if err == sql.ErrNoRows {
		summary.ProdukTerlaris = models.BestSellingProduct{
			Nama:       "-",
			QtyTerjual: 0,
		}
	} else if err != nil {
		return summary, err
	}

	return summary, nil
}
