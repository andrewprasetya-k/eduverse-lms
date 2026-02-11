# 🏫 School Management System - Backend To-Do List

Daftar tugas pengembangan modul sekolah dan fitur pendukungnya (Superadmin, Admin, Guru, & Murid).

---

## 🟢 Phase 1: Foundations (Current Progress)
*Fokus pada pembersihan CRUD dan validasi data.*

- [ ] **Fix DTO Validation**: Tambahkan `binding:"required"` pada struct `CreateSchoolDTO`.
- [ ] **Data Sanitization**: Implementasi `strings.TrimSpace()` di Service agar tidak ada nama sekolah berisi spasi kosong.
- [ ] **Standardize Response**: Buat helper untuk mapping `domain.School` ke `dto.SchoolResponseDTO` agar kode Handler tidak repetitif.
- [ ] **Restore Function**: Buat endpoint `PATCH /schools/:code/restore` untuk membatalkan soft-delete.
- [ ] **Hard Delete**: Buat endpoint `DELETE /schools/:code/permanent` khusus untuk Superadmin.

---

## 🟡 Phase 2: Auth & Security (Access Control)
*Memastikan hanya orang yang berhak yang bisa mengakses data.*

- [ ] **User Model**: Buat tabel `users` dengan field `role` (Superadmin, Admin Sekolah, Guru, Murid).
- [ ] **JWT Implementation**: Buat sistem login dan generate token.
- [ ] **RBAC Middleware**:
    - [ ] `IsSuperAdmin`: Bisa CRUD semua sekolah.
    - [ ] `IsSchoolAdmin`: Hanya bisa Update/View sekolah mereka sendiri.
- [ ] **School-User Association**: Tambahkan field `school_id` pada tabel `users` untuk mengikat user ke sekolah tertentu.

---

## 🔵 Phase 3: School Ecosystem (Internal Features)
*Fitur yang dibutuhkan oleh Kepsek, Guru, dan Murid.*

- [ ] **Academic Year**: CRUD Tahun Ajaran (Contoh: 2025/2026 Ganjil).
- [ ] **Classroom Management**: CRUD Data Kelas (Contoh: X-IPA-1).
- [ ] **Subject Master**: Daftar mata pelajaran (Matematika, Biologi, dll).
- [ ] **Public Profile**: Endpoint `GET /schools/public/:code` untuk profil ringan sekolah (tanpa data sensitif).

---

## 🟣 Phase 4: Media & Assets
*Mengelola file pendukung sekolah.*

- [ ] **Upload Logo**: Integrasi sistem upload gambar (Local Storage atau Cloud S3).
- [ ] **Image Resizing**: Otomatis mengecilkan ukuran logo sekolah agar hemat bandwidth.
- [ ] **Default Logo**: Sistem fallback jika sekolah belum mengunggah logo.

---

## 🔴 Phase 5: Optimization & Maintenance
- [ ] **Pagination**: Tambahkan fitur `?page=1&limit=10` pada list sekolah.
- [ ] **Search & Filter**: Filter sekolah berdasarkan lokasi atau status aktif.
- [ ] **Audit Logs**: Catat aktivitas `Siapa mengubah data Apa` di setiap transaksi sekolah.

---

## 📖 Notes
- Gunakan format tanggal `DD-MM-YYYY` untuk response frontend.
- Selalu gunakan `Transaction` (DB Tx) jika nanti ada proses Create School yang sekaligus membuat User Admin Sekolah.