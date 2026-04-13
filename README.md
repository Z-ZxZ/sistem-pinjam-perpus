# Sistem Pinjam Perpus — Project Iseng buat Porto (Gabut Semester 2)

Halo bang! Ini project iseng gua bikin sistem perpustakaan modern. Sengaja dibikin rada over-engineered pake (Frontend Next.js & Backend Golang) biar keliatan jago aja pas ditaruh di CV atau porto wkwk. Padahal jujur aja baru kemaren belajar Golang di Youtube.

## Stack Teknologi (Biar Keliatan Keren)

- **Frontend**: Next.js 15 (App Router), Tailwind CSS (karena kelamaan pake Vanilla CSS), Framer Motion, Lucide Icons.
- **Backend**: Golang (pakai sok-sokan Clean Architecture biar HRD atau recruiter takjub), standard library `net/http`, Auth pake JWT, sama database PostgreSQL.
- **Infrastruktur**: Docker, Redis, Nginx (biar kerasa kayak anak DevOps dikit wkwk).

## Cara Ngerun (Kalo Error Googling Aja)

1. **Prasyarat**: Udah harus install Docker sama Docker Compose yak. Kalo belum, tonton dlu tutorial di Youtube.
2. **Setup Environment**:
   - Kalo mau gampang, langsung pake default ae yang ada di `docker-compose.yml`, ga usah aneh-aneh ntar error mumet.
3. **Build & Run**:
   Tinggal ketik ini aja di terminal, sambil seduh kopi nunggu build nya:

   ```bash
   docker-compose up --build
   ```

4. **Masukin Data Dummy (Seeding)**:
   Ini penting nih, masukin dlu data dummy buat pamer ke temen/recruiter biar ga kosong banget websitenya:

   ```bash
   docker exec -it perpus_backend go run cmd/seeder/main.go
   ```

5. **Akses**:
   - Frontend nya buka di sini: `http://localhost:3000`
   - Backend API nya di mari: `http://localhost:8080`

## Akun Login Admin

- **Email**: `admin@perpus.com`
- **Password**: `admin123` *(jangan diganti ya wey ntar gabisa login pas mau pamer)*

## Fitur Unggulan

- Cari buku super ngebut soalnya metadata-nya asik.
- Ada denda otomatis! Bayar denda Rp 2.000/hari telat minjem (lumayan wkwk).
- Dashboard admin ada analitiknya, jadi kelihatan pro gitu.
- Desain webnya udah mantul ngab! Pake glassmorphism ala-ala app zaman now, micro-animations pokoknya biar yang mampir ke porto ga bosen liatnya.

---
**Catatan Penting dari gua**: Sumpah ya di project ini iseng doang, murni sok serius dan elegan biar kesannya kaya *engineer* beneran wkwk.
