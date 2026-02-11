# 🏫 School Domain - Single Table Completion Checklist

Daftar fitur ini dirancang agar tabel `schools` memiliki fungsionalitas penuh (Pencarian, Validasi, Manajemen Status, dan Statistik) secara mandiri.

---

## 🛠️ 1. Repository Layer (Database Access)
*Menyediakan mesin pencari dan manipulasi data pada tabel schools.*

- [ ] **Dynamic FindAll**: Implementasi fungsi `FindAll` yang bisa menerima map filter:
    - [ ] `search`: Query `LIKE` pada kolom `sch_name` dan `sch_code`.
    - [ ] `status`: Pilihan filter `active` (deleted_at IS NULL), `deleted` (deleted_at IS NOT NULL), atau `all`.
    - [ ] `pagination`: Menggunakan `.Limit()` dan `.Offset()`.
    - [ ] `sorting`: Mengurutkan berdasarkan `created_at` atau `sch_name`.
- [ ] **CheckCodeExists**: Fungsi efisien untuk mengecek apakah `sch_code` sudah ada (untuk validasi sebelum insert/update).
- [ ] **RestoreSchool**: Fungsi untuk mengubah `deleted_at` menjadi `NULL` berdasarkan `sch_code`.
- [ ] **HardDeleteSchool**: Fungsi `Unscoped().Delete()` untuk menghapus data secara fisik dari tabel.

---

## 🧠 2. Service Layer (Business Logic)
*Logika pemrosesan data sebelum dikirim atau disimpan.*

- [ ] **Input Sanitization**: Fungsi otomatis untuk `TrimSpace` pada nama, alamat, dan email agar tidak ada data "spasi kosong".
- [ ] **Automatic Code Generator**: (Jika belum) Logika untuk membuat `sch_code` otomatis jika user tidak mengisinya.
- [ ] **Validation Logic**: 
    - [ ] Cek duplikasi email sekolah.
    - [ ] Cek duplikasi nomor telepon sekolah.
- [ ] **School Statistics**: Fungsi untuk menghitung total baris:
    - [ ] `CountAll()`
    - [ ] `CountActive()`
    - [ ] `CountDeleted()`

---

## 🌐 3. Handler Layer (API Endpoints)
*Menyediakan akses bagi user/frontend ke fitur-fitur di atas.*

- [ ] **Unified GET /schools**: Mengganti semua fungsi GET list menjadi satu endpoint yang menerima Query Params:
    - `?search=...&status=...&page=...&limit=...`
- [ ] **PATCH /schools/:code/restore**: Endpoint khusus untuk mengaktifkan kembali sekolah.
- [ ] **DELETE /schools/:code/permanent**: Endpoint untuk menghapus data secara permanen.
- [ ] **GET /schools/check-code/:code**: Endpoint cepat untuk validasi ketersediaan kode di frontend.
- [ ] **GET /schools/summary**: Endpoint yang mengembalikan jumlah total (Statistik).

---

## 🛡️ 4. Data Integrity (GORM/DB Level)
- [ ] **Constraint Handling**: Menangani pesan error database agar lebih ramah (misal: "Email sudah terdaftar" daripada "Error 1062 Duplicate Entry").
- [ ] **Default Values**: Memastikan `sch_logo` bisa menerima `NULL` jika sekolah belum memiliki logo.