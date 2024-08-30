# CRUD Simpel

CRUD Simpel adalah aplikasi sederhana yang dikembangkan menggunakan bahasa pemrograman Go untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data pengguna, produk, dan transaksi. Aplikasi ini menggunakan pendekatan sharding untuk distribusi data dan menangani concurrency dengan mekanisme mutex baca-tulis.

## Struktur Proyek

- **`datastore/`**: Berisi implementasi penyimpanan data menggunakan peta (map) untuk simulasi database dengan pendekatan sharding untuk distribusi data.
- **`models/`**: Berisi struktur data (model) yang digunakan dalam aplikasi.
- **`response/`**: Berisi fungsi untuk menangani response API, termasuk response error dan response sukses.
- **`validator/`**: Berisi fungsi untuk validasi data input.
- **`handlers/`**: Berisi handler untuk endpoint API untuk pengguna, produk, dan transaksi.

## Sharding

Proyek ini menggunakan pendekatan sharding untuk mendistribusikan data pada penyimpanan dengan membagi data ke dalam beberapa shard. Setiap shard memiliki mutex baca-tulis dan variabel kondisi untuk manajemen akses paralel. Berikut adalah rincian implementasinya:

- **`shardCount`**: Menentukan jumlah shard, diatur ke 3 dalam implementasi ini.
- **`Shard`**: Struktur yang mewakili sebuah shard dengan mutex baca-tulis dan variabel kondisi.
- **`UserShards`**: Array yang menyimpan shard untuk data pengguna.
- **`ProductShards`**: Array yang menyimpan shard untuk data produk.
- **`TransactionShards`**: Array yang menyimpan shard untuk data transaksi.

## Instalasi

1. **Clone Repositori:**

    ```bash
    git clone https://github.com/ekokurniawann/simpel-crud-go.git
    cd simpel-crud-go
    ```

2. **Instal Dependensi:** Pastikan Anda telah menginstal Go di sistem Anda. Instal dependensi menggunakan Go modules:

    ```bash
    go mod tidy
    ```

## Menjalankan Aplikasi

1. **Jalankan Server:** Untuk menjalankan server, gunakan perintah berikut:

    ```bash
    go run main.go
    ```

    Server akan berjalan pada `http://localhost:8080` secara default.

## Endpoints API

### Users

- **POST /users**
    - Membuat pengguna baru.
    - **Request Body:** `{"ID": "string", "Name": "string", "Email": "string"}`
    - **Response:** `201 Created` dan data pengguna yang dibuat.

- **GET /users/{id}**
    - Mengambil data pengguna berdasarkan ID.
    - **Response:** `200 OK` dan data pengguna jika ditemukan, `404 Not Found` jika tidak ditemukan.

- **PUT /users/{id}**
    - Memperbarui data pengguna berdasarkan ID.
    - **Request Body:** `{"Name": "string", "Email": "string"}`
    - **Response:** `200 OK` dan data pengguna yang diperbarui, `404 Not Found` jika tidak ditemukan.

- **DELETE /users/{id}**
    - Menghapus pengguna berdasarkan ID.
    - **Response:** `204 No Content` jika berhasil, `404 Not Found` jika tidak ditemukan.

### Products

- **POST /products**
    - Membuat produk baru.
    - **Request Body:** `{"ID": "string", "Name": "string", "Description": "string", "Price": float}`
    - **Response:** `201 Created` dan data produk yang dibuat.

- **GET /products/{id}**
    - Mengambil data produk berdasarkan ID.
    - **Response:** `200 OK` dan data produk jika ditemukan, `404 Not Found` jika tidak ditemukan.

- **PUT /products/{id}**
    - Memperbarui data produk berdasarkan ID.
    - **Request Body:** `{"Name": "string", "Description": "string", "Price": float}`
    - **Response:** `200 OK` dan data produk yang diperbarui, `404 Not Found` jika tidak ditemukan.

- **DELETE /products/{id}**
    - Menghapus produk berdasarkan ID.
    - **Response:** `204 No Content` jika berhasil, `404 Not Found` jika tidak ditemukan.

### Transactions

- **POST /transactions**
    - Membuat transaksi baru.
    - **Request Body:** `{"ID": "string", "UserID": "string", "ProductID": "string", "Quantity": int, "Total": float}`
    - **Response:** `201 Created` dan data transaksi yang dibuat.

- **GET /transactions/{id}**
    - Mengambil data transaksi berdasarkan ID.
    - **Response:** `200 OK` dan data transaksi jika ditemukan, `404 Not Found` jika tidak ditemukan.

- **PUT /transactions/{id}**
    - Memperbarui data transaksi berdasarkan ID.
    - **Request Body:** `{"UserID": "string", "ProductID": "string", "Quantity": int, "Total": float}`
    - **Response:** `200 OK` dan data transaksi yang diperbarui, `404 Not Found` jika tidak ditemukan.

- **DELETE /transactions/{id}**
    - Menghapus transaksi berdasarkan ID.
    - **Response:** `204 No Content` jika berhasil, `404 Not Found` jika tidak ditemukan.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan buka [issues](https://github.com/ekokurniawann/simpel-crud-go/issues) untuk melihat masalah yang ada atau buat issue baru. Kirim pull request untuk perubahan yang Anda buat.
