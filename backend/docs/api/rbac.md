# 🔐 RBAC (Role-Based Access Control) API Documentation

Base URL: `/api/rbac`

---

## Overview

Role-Based Access Control (RBAC) mengamankan API endpoints berdasarkan role user di setiap school. Sistem mendukung multi-school dengan role berbeda per school.

### Roles

| Role | Permissions |
|------|-------------|
| `super_admin` | Full access ke semua sekolah dan fitur |
| `admin` | Manage sekolah tertentu (academic years, terms, users, subjects) |
| `teacher` | Manage kelas, materials, assignments yang diajar |
| `student` | Akses read-only + submit assignment |

---

## 1. Role Management

### List All Roles
- **URL:** `/roles`
- **Method:** `GET`
- **Auth:** Required

**Response Example:**
```json
[
  {
    "roleId": "uuid",
    "roleName": "super_admin",
    "createdAt": "02-01-2006 15:04:05"
  },
  {
    "roleId": "uuid",
    "roleName": "teacher",
    "createdAt": "02-01-2006 15:04:05"
  }
]
```

### Create Role
- **URL:** `/roles`
- **Method:** `POST`
- **Auth:** Required (super_admin only)
- **Body:**
```json
{
  "roleName": "teacher"
}
```

### Get Role by ID
- **URL:** `/roles/:id`
- **Method:** `GET`
- **Auth:** Required

### Update Role Name
- **URL:** `/roles/:id`
- **Method:** `PATCH`
- **Auth:** Required (super_admin only)
- **Body:**
```json
{
  "roleName": "senior_teacher"
}
```

### Delete Role
- **URL:** `/roles/:id`
- **Method:** `DELETE`
- **Auth:** Required (super_admin only)

---

## 2. User Role Management (Assignments)

Assigning roles to users within a school context.

### Assign Role to User
- **URL:** `/user-roles`
- **Method:** `POST`
- **Auth:** Required (admin, super_admin)
- **Body:**
```json
{
  "schoolUserId": "uuid",
  "roleId": "uuid"
}
```

### Remove Role from User
- **URL:** `/user-roles?schoolUserId=...&roleId=...`
- **Method:** `DELETE`
- **Auth:** Required (admin, super_admin)

### List User's Roles
- **URL:** `/user-roles/:schoolUserId`
- **Method:** `GET`
- **Auth:** Required

**Response Example:**
```json
[
  {
    "urol_id": "uuid",
    "urol_scu_id": "uuid",
    "urol_rol_id": "uuid",
    "role": {
      "rol_id": "uuid",
      "rol_name": "teacher"
    }
  }
]
```

### Update User Roles (Sync)
Replace all roles for a user.
- **URL:** `/user-roles/:schoolUserId`
- **Method:** `PATCH`
- **Auth:** Required (admin, super_admin)
- **Body:**
```json
{
  "roleIds": ["role-uuid-1", "role-uuid-2"]
}
```

---

## 3. RBAC Middleware

### School Context Header

**All protected endpoints require school context via:**

**Priority 1: Header (Recommended)**
```
X-School-ID: uuid-school-id
```

**Priority 2: URL Parameter (Fallback)**
```
/api/schools/:schoolCode/...
```

### Request Example

```bash
POST /api/classes
Authorization: Bearer <token>
X-School-ID: uuid-school-id
Content-Type: application/json

{
  "cls_code": "12-IPA-1",
  "cls_title": "Kelas 12 IPA 1"
}
```

---

## 4. Protected Endpoints

### School Management
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/schools` | POST | super_admin |
| `/schools/:schoolCode` | PATCH | admin, super_admin |
| `/schools/:schoolCode` | DELETE | admin, super_admin |

### Academic Structure
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/academic-years` | POST | admin, super_admin |
| `/terms` | POST | admin, super_admin |
| `/subjects` | POST | admin, super_admin |

### Class Management
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/classes` | POST | admin, teacher |
| `/classes/:id` | PATCH | admin, teacher |
| `/classes/:id` | DELETE | admin |

### Learning Content
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/materials` | POST | teacher |
| `/assignments` | POST | teacher |
| `/assignments/submit/:id` | POST | student |
| `/assignments/assess/:id` | POST | teacher |

### User Management
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/users` | POST | admin, super_admin |
| `/users` | GET | admin, super_admin |
| `/users/:id` | DELETE | admin, super_admin |

### Enrollment
| Endpoint | Method | Required Role |
|----------|--------|---------------|
| `/enrollments` | POST | admin, teacher |
| `/enrollments/:id` | DELETE | admin, teacher |

---

## 5. Error Responses

### 400 Bad Request
School context tidak ditemukan.

```json
{
  "error": "School context required (X-School-ID header or schoolCode param)"
}
```

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

---

## 6. Setup & Testing

### Initial Setup

1. **Create Roles** (via API or SQL)
```bash
POST /api/rbac/roles
{
  "roleName": "super_admin"
}
```

2. **Enroll User to School**
```bash
POST /api/school-users/enroll
{
  "scu_usr_id": "user-uuid",
  "scu_sch_id": "school-uuid"
}
```

3. **Assign Role**
```bash
POST /api/rbac/user-roles
{
  "schoolUserId": "school-user-uuid",
  "roleId": "role-uuid"
}
```

### Frontend Integration

```javascript
// 1. User login
const { token } = await login(email, password);

// 2. Get user's schools
const schools = await fetch(`/api/school-users/user/${userId}`, {
  headers: { 'Authorization': `Bearer ${token}` }
});

// 3. User selects a school
const selectedSchoolId = schools[0].school.sch_id;

// 4. Make requests with X-School-ID header
const response = await fetch('/api/classes', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`,
    'X-School-ID': selectedSchoolId,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    cls_code: '12-IPA-1',
    cls_title: 'Kelas 12 IPA 1'
  })
});
```

### Test Examples

**Test as Admin (Success)**
```bash
POST /api/classes
Authorization: Bearer <admin-token>
X-School-ID: <school-id>

{
  "cls_code": "12-IPA-1",
  "cls_title": "Kelas 12 IPA 1"
}

# Response: 200 OK
```

**Test as Student (Fail)**
```bash
POST /api/classes
Authorization: Bearer <student-token>
X-School-ID: <school-id>

{
  "cls_code": "12-IPA-1",
  "cls_title": "Kelas 12 IPA 1"
}

# Response: 403 Forbidden
{
  "error": "Forbidden: insufficient permissions"
}
```

---

## 7. Implementation Notes

### Multi-School Support
- User dapat memiliki role berbeda di sekolah berbeda
- Frontend mengirim `X-School-ID` header untuk specify context
- Middleware otomatis validate membership dan role

### Backward Compatible
- Endpoint dengan `schoolCode` di URL tetap work
- Tidak ada breaking changes pada API contract
- Handler code tidak perlu diubah

### Security Features
- Cross-tenant isolation (user tidak bisa akses school lain)
- Role-based permissions (action restricted by role)
- Fail-secure (default deny jika tidak ada role match)

### Future Enhancements
- [ ] Permission-based access (granular control)
- [ ] Resource ownership check (creator-only modifications)
- [ ] Class-level access (teacher/student specific to class)
- [ ] Audit logging untuk access attempts
