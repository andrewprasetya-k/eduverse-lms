# Eduverse LMS - Backend TODO List

## Arsitektur & Perbaikan (Urgent)
- [ ] **Fix Subject Search Logic**: Update `GetSubjectByCode` di Repository, Service, dan Handler agar selalu menyertakan `school_id`. Saat ini masih mencari global berdasarkan kode saja, yang berisiko tertukar antar sekolah.

## Fitur School
- [ ] Implementasi validasi `LogoID` (cek apakah media ID benar-benar ada).

## Fitur Subject
- [ ] Lanjutkan pembuatan `SubjectHandler` (Update & Delete).
- [ ] Registrasi route Subject di `main.go`.

## Fitur Selanjutnya
- [ ] Entitas `Users` & `Academic Years`.
- [ ] Implementasi Auth (JWT).
- [ ] Implementasi Role-based Access Control (RBAC).
