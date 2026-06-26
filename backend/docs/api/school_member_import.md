# School Member Import API

MVP import warga sekolah digunakan oleh Admin Sekolah untuk menambahkan siswa,
guru, atau admin sekolah ke sekolah aktif melalui file CSV.

Import ini tidak membuka daftar akun global lintas platform untuk Admin Sekolah.
Jika email sudah ada sebagai akun global aktif, akun tersebut dipakai ulang dan
ditautkan ke sekolah aktif.

## Preview Import

- **URL:** `/admin/school-members/import/preview`
- **Method:** `POST`
- **Auth:** Admin sekolah pada sekolah aktif
- **Content-Type:** `multipart/form-data`
- **Fields:**
  - `file`: CSV dengan kolom `fullName,email,role,classCode`

`classCode` bersifat opsional dan hanya berlaku untuk role `student`.

Response:

```json
{
  "rows": [
    {
      "rowNumber": 2,
      "fullName": "Budi Santoso",
      "email": "budi@siswa.sch.id",
      "role": "student",
      "classCode": "X-IPA-1",
      "status": "valid",
      "errors": []
    }
  ],
  "validCount": 1,
  "invalidCount": 0
}
```

## Commit Import

- **URL:** `/admin/school-members/import/commit`
- **Method:** `POST`
- **Auth:** Admin sekolah pada sekolah aktif

Request:

```json
{
  "defaultPassword": "InitialPassword123!",
  "rows": [
    {
      "rowNumber": 2,
      "fullName": "Budi Santoso",
      "email": "budi@siswa.sch.id",
      "role": "student",
      "classCode": "X-IPA-1"
    }
  ]
}
```

Behavior:

- All-or-nothing commit.
- `fullName`, `email`, dan `role` wajib ada.
- `role` hanya boleh `student`, `teacher`, atau `admin`.
- `super_admin` selalu ditolak.
- Email duplikat dalam file ditolak.
- `defaultPassword` wajib diisi dan hanya dipakai untuk akun baru.
- Jika email sudah ada sebagai akun global aktif, user dipakai ulang.
- Membership `school_users` dibuat untuk sekolah aktif jika belum ada.
- Role dimasukkan ke `user_roles` tanpa menghapus role lain.
- Jika `classCode` diisi dan role adalah `student`, student dienroll ke kelas
  aktif tersebut.
- Teacher class assignment dan subject assignment tidak dilakukan oleh import ini.

Response:

```json
{
  "importedCount": 1,
  "skippedCount": 0,
  "failedCount": 0,
  "results": [
    {
      "rowNumber": 2,
      "fullName": "Budi Santoso",
      "email": "budi@siswa.sch.id",
      "role": "student",
      "classCode": "X-IPA-1",
      "status": "imported"
    }
  ]
}
```
