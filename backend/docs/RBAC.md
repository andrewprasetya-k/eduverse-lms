# 🔐 RBAC Implementation Guide

## Overview

Role-Based Access Control (RBAC) telah diimplementasikan untuk mengamankan API endpoints berdasarkan role user di setiap school.

## Roles

Sistem mendukung 4 role utama:

1. **super_admin** - Full access ke semua sekolah dan fitur
2. **admin** - Manage sekolah tertentu (academic years, terms, users, subjects)
3. **teacher** - Manage kelas, materials, assignments yang diajar
4. **student** - Akses read-only + submit assignment

## Middleware

### 1. `RequireSchoolAccess(schoolService)`

Memastikan user terdaftar di school yang diakses.

**Usage:**
```go
schoolAPI.GET("/:schoolCode", 
    middleware.RequireSchoolAccess(schoolService), 
    handler.GetSchool)
```

**Behavior:**
- Extract `schoolCode` dari URL param
- Convert code ke ID via `schoolService.ConvertCodeToID()`
- Check apakah user ada di `school_users` table
- Return `403 Forbidden` jika bukan member

### 2. `RequireRole(schoolService, ...roles)`

Memastikan user memiliki salah satu role yang diizinkan.

**Usage:**
```go
schoolAPI.POST("/", 
    middleware.RequireRole(schoolService, "admin", "super_admin"), 
    handler.CreateSchool)
```

**Behavior:**
- Get user roles dari `user_roles` table
- Check apakah user punya salah satu role yang diizinkan
- Return `403 Forbidden` jika tidak punya role yang sesuai

## Protected Endpoints

### School Management
- `POST /schools` - super_admin only
- `PATCH /schools/:schoolCode` - admin, super_admin (+ school access)
- `DELETE /schools/:schoolCode` - admin, super_admin (+ school access)

### Academic Structure
- `POST /academic-years` - admin, super_admin
- `POST /terms` - admin, super_admin
- `POST /subjects` - admin, super_admin

### Class Management
- `POST /classes` - admin, teacher
- `PATCH /classes/:id` - admin, teacher
- `DELETE /classes/:id` - admin only

### Learning Content
- `POST /materials` - teacher only
- `POST /assignments` - teacher only
- `POST /assignments/submit/:id` - student only
- `POST /assignments/assess/:id` - teacher only

### User Management
- `POST /users` - admin, super_admin
- `GET /users` - admin, super_admin
- `DELETE /users/:id` - admin, super_admin

### Enrollment
- `POST /enrollments` - admin, teacher
- `DELETE /enrollments/:id` - admin, teacher

## Error Responses

### 401 Unauthorized
User tidak login atau token invalid.

```json
{
  "error": "Unauthorized"
}
```

### 403 Forbidden - Insufficient Permissions
User tidak punya role yang sesuai.

```json
{
  "error": "Forbidden: insufficient permissions"
}
```

### 403 Forbidden - Not School Member
User bukan member dari school yang diakses.

```json
{
  "error": "Forbidden: not a member of this school"
}
```

## Setup Initial Roles

Untuk setup roles awal, jalankan query berikut:

```sql
INSERT INTO roles (rol_id, rol_name, created_at) VALUES
  (gen_random_uuid(), 'super_admin', now()),
  (gen_random_uuid(), 'admin', now()),
  (gen_random_uuid(), 'teacher', now()),
  (gen_random_uuid(), 'student', now());
```

## Assign Role ke User

1. User harus terdaftar di school via `school_users`
2. Assign role via endpoint:

```bash
POST /api/rbac/user-roles
{
  "urol_scu_id": "school_user_id",
  "urol_rol_id": "role_id"
}
```

## Testing

### Test sebagai Admin
```bash
# Login sebagai admin
POST /api/login
{
  "email": "admin@school.com",
  "password": "password"
}

# Gunakan token untuk create class
POST /api/classes
Authorization: Bearer <token>
{
  "cls_code": "12-IPA-1",
  ...
}
```

### Test sebagai Student (Should Fail)
```bash
# Login sebagai student
POST /api/login
{
  "email": "student@school.com",
  "password": "password"
}

# Try create class (akan dapat 403)
POST /api/classes
Authorization: Bearer <token>
{
  "cls_code": "12-IPA-1",
  ...
}

# Response:
{
  "error": "Forbidden: insufficient permissions"
}
```

## Implementation Details

### Repository Layer
File: `internal/repository/rbac_repo.go`

Helper methods:
- `GetUserRoleNamesInSchool(userID, schoolID)` - Get role names
- `IsUserInSchool(userID, schoolID)` - Check membership
- `GetSchoolUserID(userID, schoolID)` - Get school_user ID

### Middleware Layer
File: `internal/middleware/rbac_middleware.go`

- `InitRBAC(repo)` - Initialize dengan repository
- `RequireSchoolAccess(schoolService)` - School membership check
- `RequireRole(schoolService, ...roles)` - Role-based check

### Route Registration
File: `cmd/api/main.go`

Middleware ditambahkan di route registration tanpa mengubah handler code.

## Notes

- Middleware bersifat **additive** - tidak mengubah endpoint URL atau payload
- Backward compatible - handler code tidak perlu diubah
- Multi-school aware - user bisa punya role berbeda di school berbeda
- Chainable - bisa combine `RequireSchoolAccess` + `RequireRole`

## Future Enhancements

- [ ] Permission-based access (granular control)
- [ ] Resource ownership check (creator-only modifications)
- [ ] Class-level access (teacher/student specific to class)
- [ ] Audit logging untuk access attempts
