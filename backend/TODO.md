# 🎓 Eduverse LMS - Development Progress

1. buat header yg diambil dari response json. endpoint yang berupa get by school atau get by selain id harus ada header sekolahnya
2. buat endpoint untuk liat users/members yang ada di enrollment tertentu (jangan lupa header)
3. ubah error binding (ga ada foreign key, dan lain lain) jadi lebih tersirat biar ga tampilin nama kolom pure di UI
4. sambungin materials ke class, lalu kasih headernya itu berupa nama subjectnya

## 🚀 Next Priorities (Future Tasks)

- [ ] **Authentication Middleware**: Implement JWT/Session protection for all routes.
- [ ] **File Upload Integration**: Connect `medias` module to actual S3/Supabase storage.
- [ ] **Advanced Grading**: Implementation of Assessment Weights logic.
- [ ] **Frontend Integration**: Start building the UI based on `docs/api/`.
- [ ] **Email Service**: Implement actual email sending for confirmation and password reset.
