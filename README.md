# Sistem Pinjam Perpus — Production Ready

Sistem manajemen perpustakaan modern dengan arsitektur terpisah (Frontend Next.js & Backend Golang) yang dibangun berdasarkan blueprint produksi.

## Teknologi Utama

- **Frontend**: Next.js 15 (App Router), Tailwind CSS, Framer Motion, Lucide Icons.
- **Backend**: Golang (Clean Architecture), standard library `net/http`, JWT Auth, PostgreSQL.
- **Infrastruktur**: Docker, Redis, Nginx (ready).

## Cara Menjalankan (Production Setup)

1. **Prasyarat**: Pastikan Docker dan Docker Compose sudah terinstal.
2. **Setup Environment**: 
   - Salin file environment (jika ada) atau biarkan default di `docker-compose.yml`.
3. **Build & Run**:
   ```bash
   docker-compose up --build
   ```
4. **Seeding Data**:
   Setelah sistem berjalan, jalankan seeder untuk mengisi data awal (Admin & Buku):
   ```bash
   docker exec -it perpus_backend go run cmd/seeder/main.go
   ```
5. **Akses**:
   - Frontend: `http://localhost:3000`
   - Backend API: `http://localhost:8080`

## Akun Demo (Admin)
- **Email**: `admin@perpus.com`
- **Password**: `admin123`

## Fitur Ungkulan
- Pencarian buku cepat dengan integrasi metadata global.
- Sistem denda otomatis (Rp 2000/hari keterlambatan).
- Dashboard analitik untuk Admin dan riwayat untuk Member.
- Desain premium dengan micro-animations dan glassmorphism.

---
**Catatan Penting**: Mengikuti instruksi blueprint, sistem ini tidak menggunakan emoji sama sekali dalam seluruh antarmuka dan basis kode.
