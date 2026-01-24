package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Struct data Produk dan Category

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type Category struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Description string `json:"description"`
}

// Inisialisasi variabel data Produk dan Category

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "Kecap", Harga: 12000, Stok: 20},
}

var category = []Category{
	{ID: 1, Nama: "Makanan", Description: "Kategori untuk semua makanan"},
	{ID: 2, Nama: "Minuman", Description: "Kategori untuk semua minuman"},
	{ID: 3, Nama: "Bumbu", Description: "Kategori untuk semua bumbu dapur"},
}

// Deklarasi Fungsi untuk menangani request GET, PUT, DELETE berdasarkan ID Produk

func getProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
	//get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	//ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	//get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Produk berhasil dihapus",
			})
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// Deklarasi Fungsi untuk menangani request GET, PUT, DELETE berdasarkan Category

// Helper untuk mengambil ID dari URL
func getIDFromURL(r *http.Request, prefix string) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}

// Handler untuk menangani banyak data
func handleCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(category) //Mengambil semua data category

	case "POST":
		var newCategory Category
		if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil { //Baca request body
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Logic penambahan ID
		newCategory.ID = len(category) + 1
		category = append(category, newCategory) //Menambahkan data category baru

		// Status 201 Created
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCategory)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler untuk menangani single data berdasarkan ID
func handleCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := getIDFromURL(r, "/api/category/") //Memanggil helper untuk mendapatkan ID dari URL
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		for _, c := range category {
			if c.ID == id {
				json.NewEncoder(w).Encode(c) //Mengembalikan data category sesuai ID
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)
	case "PUT":
		var updatedCategory Category
		if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil { //Baca request body
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		for i := range category {
			if category[i].ID == id {
				updatedCategory.ID = id
				category[i] = updatedCategory //Update data category sesuai ID
				json.NewEncoder(w).Encode(updatedCategory)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)
	case "DELETE":
		for i, c := range category {
			if c.ID == id {
				category = append(category[:i], category[i+1:]...)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Category berhasil dihapus",
				})
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Deklarasi Fungsi utama untuk menjalankan server

func main() {
	// GET localhost:8040/api/produk/{id}
	// PUT localhost:8040/api/produk/{id}
	// DELETE localhost:8040/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// GET All Produk & POST localhost:8040/produk
	http.HandleFunc("/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// Menambahkan produk baru ke variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8040/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API Running",
		})
	})
	fmt.Println("Server running on port 8040")

	// Routing untuk Category
	http.HandleFunc("/api/category", handleCategories)

	http.HandleFunc("/api/category/", handleCategoryByID)

	err := http.ListenAndServe(":8040", nil)
	if err != nil {
		fmt.Println("Gagal running server:", err)
	}
}
