package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/databases"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

// Deklarasi Fungsi utama untuk menjalankan server
func main() {
	//Config viper untuk membaca .env
	viper.AutomaticEnv() // Membaca variabel dari .env secara otomatis
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup Koneksi Database
	db, err := databases.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Inisialisasi Repository, Service, dan Handler untuk Produk
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Routing untuk Produk
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	// Inisialisasi Repository, Service, dan Handler untuk Kategori
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Routing untuk Kategori
	http.HandleFunc("/api/category", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Routing untuk Transaction
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST

	// Report
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	// Routing untuk Report
	http.HandleFunc("/api/report", reportHandler.HandleReport)

	// localhost:8040/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API Running",
		})
	})
	fmt.Println("Server running on port 8040")

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
