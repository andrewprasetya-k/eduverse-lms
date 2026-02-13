# 🏫 School Domain - Single Table Completion Checklist

Daftar fitur ini dirancang agar tabel `schools` memiliki fungsionalitas penuh (Pencarian, Validasi, Manajemen Status, dan Statistik) secara mandiri.

---

## 🛠️ 1. Repository Layer (Database Access)

_Menyediakan mesin pencari dan manipulasi data pada tabel schools._

- [x] Ganti schoolCode ke schoolID untuk tabel/entitas schools
- [x] **Dynamic FindAll**: Implementasi fungsi `FindAll` yang bisa menerima map filter:
  - [x] `search`: Query `LIKE` pada kolom `sch_name` dan `sch_code`.
  - [x] `status`: Pilihan filter `active` (deleted_at IS NULL), `deleted` (deleted_at IS NOT NULL), atau `all`.
  - [x] `pagination`: Menggunakan `.Limit()` dan `.Offset()`.
  - [x] `sorting`: Mengurutkan berdasarkan `created_at` atau `sch_name`.
- [x] **CheckCodeExists**: Fungsi efisien untuk mengecek apakah `sch_code` sudah ada (untuk validasi sebelum insert/update).
- [x] **RestoreSchool**: Fungsi untuk mengubah `deleted_at` menjadi `NULL` berdasarkan `sch_code`.
- [x] **HardDeleteSchool**: Fungsi `Unscoped().Delete()` untuk menghapus data secara fisik dari tabel.

---

## 🧠 2. Service Layer (Business Logic)

_Logika pemrosesan data sebelum dikirim atau disimpan._

- [x] **Input Sanitization**: Fungsi otomatis untuk `TrimSpace` pada nama, alamat, dan email agar tidak ada data "spasi kosong".
- [x] **Automatic Code Generator**: (Jika belum) Logika untuk membuat `sch_code` otomatis jika user tidak mengisinya.
- [x] **Validation Logic**:
  - [x] Cek duplikasi email sekolah.
  - [x] Cek duplikasi nomor telepon sekolah.
- [x] **School Statistics**: Fungsi untuk menghitung total baris:
  - [x] `CountAll()`
  - [x] `CountActive()`
  - [x] `CountDeleted()`

---

## 🌐 3. Handler Layer (API Endpoints)

_Menyediakan akses bagi user/frontend ke fitur-fitur di atas._

- [x] **Unified GET /schools**: Mengganti semua fungsi GET list menjadi satu endpoint yang menerima Query Params:
  - `?search=...&status=...&page=...&limit=...`
- [x] **PATCH /schools/:code/restore**: Endpoint khusus untuk mengaktifkan kembali sekolah.
- [x] **DELETE /schools/:code/permanent**: Endpoint untuk menghapus data secara permanen.
- [x] **GET /schools/check-code/:code**: Endpoint cepat untuk validasi ketersediaan kode di frontend.
- [x] **GET /schools/summary**: Endpoint yang mengembalikan jumlah total (Statistik).

---

## 🛡️ 4. Data Integrity (GORM/DB Level)

- [x] **Constraint Handling**: Menangani pesan error database agar lebih ramah (misal: "Email sudah terdaftar" daripada "Error 1062 Duplicate Entry").
- [x] **Default Values**: Memastikan `sch_logo` bisa menerima `NULL` jika sekolah belum memiliki logo.

---

## 🔒 5. User & Auth TODOs
- [ ] **Email Confirmation**: Add email confirmation flow before allowing password changes or critical updates.
- [x] **Password Hashing**: Implement secure password hashing using bcrypt.
- [x] **Global User Management**: CRUD operations for users.
- [x] **School Membership**: Connect users to specific schools.
