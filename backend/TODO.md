# 🎓 Eduverse LMS - Development Progress

1. buat header yg diambil dari response json. endpoint yang berupa get by school atau get by selain id harus ada header sekolahnya
2. cek bagaimana cara mencari permissions yang diperbolehkan sebelum melakukan operasi (apakah menggunakan tabel tambahan atau gimana)
3. ubah error binding (ga ada foreign key, dan lain lain) jadi lebih tersirat biar ga tampilin nama kolom pure di UI
4. kenapa ga kasih kolom role juga di subject_class sama kayak di tabel enrollments buat tandain siapa guru dan siapa murid?
5. pusing pusing pusing

## 🚀 Next Priorities (Future Tasks)

- [ ] **Authentication Middleware**: Implement JWT/Session protection for all routes.
- [ ] **File Upload Integration**: Connect `medias` module to actual S3/Supabase storage.
- [ ] **Advanced Grading**: Implementation of Assessment Weights logic.
- [ ] **Frontend Integration**: Start building the UI based on `docs/api/`.
- [ ] **Email Service**: Implement actual email sending for confirmation and password reset.
