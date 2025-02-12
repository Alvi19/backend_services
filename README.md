Backend Services

📌 Deskripsi

Backend Services adalah proyek backend yang dikembangkan menggunakan Golang dengan framework Fiber. Proyek ini menggunakan PostgreSQL sebagai database utama dan diimplementasikan dengan arsitektur microservices.

🚀 Teknologi yang Digunakan

Golang (Fiber Framework)

PostgreSQL (Database)

JWT (JSON Web Token Authentication)

Swagger (API Documentation)

⚙️ Instalasi & Menjalankan Proyek

1. Clone Repository

git clone https://github.com/Alvi19/backend_services.git
cd backend_services

2. Buat File Konfigurasi

Salin file .env.example menjadi .env dan sesuaikan konfigurasi database serta variabel lingkungan lainnya.

cp .env.example .env

3. Instal Dependensi

go mod tidy

4. Jalankan Aplikasi

go run main.go

🔑 Autentikasi

Gunakan endpoint /api/auth/login untuk mendapatkan token JWT. Token ini diperlukan untuk mengakses endpoint yang membutuhkan autentikasi.

📖 Dokumentasi API

Swagger telah diintegrasikan dalam proyek ini. Untuk melihat dokumentasi API, jalankan aplikasi dan akses:
