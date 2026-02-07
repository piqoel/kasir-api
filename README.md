Simple Kasir API
API sederhana untuk manajemen Produk dan Kategori pada sistem kasir dibangun menggunakan Golang dengan menggunakan Layered Architecture dan terkoneksi dengan database berbasis postgresql menggunakan Supabase.

Base URL:
https://kasir-api-production-edd3.up.railway.app

Cara Tes API
Kamu bisa menggunakan aplikasi seperti Postman untuk mencoba endpoint di bawah ini.

1. Health Check
Mengecek apakah server berjalan dengan normal.

URL: /health
Method: GET

Response:
JSON
{
  "status": "ok",
  "message": "API Running"
}

2. Kategori (Categories)

🟢 Ambil Semua Kategori
URL: /category
Method: GET

🟢 Detail Satu Produk
URL: /category/{id} (Contoh: /category/1)
Method: GET

🔵 Tambah Kategori Baru
URL: /category
Method: POST

Body (JSON):
JSON
{
  "name": "Minuman",
  "description": "Segala jenis minuman dingin dan hangat"
}

🟡 Update Kategori
URL: /category/{id} (Contoh: /category/1)
Method: PUT

Body (JSON):
JSON
{
  "name": "Minuman Segar",
  "description": "Minuman dingin saja"
}

🔴 Hapus Kategori
URL: /category/{id} (Contoh: /category/1)
Method: DELETE

3. Produk (Products)

🟢 Ambil Semua Produk
URL: /produk
Method: GET

🔵 Tambah Produk Baru
URL: /produk
Method: POST

Body (JSON):
JSON
{
  "nama": "Teh Botol",
  "harga": 5000,
  "stok": 100
}

🟡 Detail Satu Produk
URL: /produk/{id} (Contoh: /produk/1)
Method: GET

🟠 Update Produk
URL: /produk/{id} (Contoh: /produk/1)
Method: PUT

Body (JSON):
JSON
{
  "nama": "Teh Botol Sosro",
  "harga": 6000,
  "stok": 90
}

🔴 Hapus Produk
URL: /produk/{id} (Contoh: /produk/1)
Method: DELETE

4. Transaksi (Transactions)

🔵 Checkout (Membuat Transaksi Baru)
Mencatat transaksi pembelian baru, mengurangi stok produk, dan menyimpan detail harga.

URL: /api/checkout
Method: POST

URL: /api/checkout
Method: POST

Body (JSON):
```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 3,
      "quantity": 1
    }
  ]
}
```

5. Laporan (Reports)

🟢 Laporan Penjualan (Sales Summary) Melihat ringkasan pendapatan, total transaksi, dan produk terlaris.

URL: /api/report Method: GET

Parameter (Query Params):

start_date (Optional): Format YYYY-MM-DD

end_date (Optional): Format YYYY-MM-DD

Contoh Penggunaan:

Laporan Hari Ini: GET /api/report

Laporan Berdasarkan Tanggal: GET /api/report?start_date=2026-01-01&end_date=2026-02-28
