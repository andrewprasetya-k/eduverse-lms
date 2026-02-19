# 🔐 RBAC (Pure Role-Based Access Control) Module API Documentation

Base URL: `/api/rbac`

## 1. Role Management
Roles are now global and not tied to specific schools.

### List All Roles
- **URL:** `/roles`
- **Method:** `GET`

**Response Example:**
```json
[
  {
    "roleId": "uuid",
    "roleName": "SUPER_ADMIN",
    "createdAt": "02-01-2006 15:04:05"
  },
  {
    "roleId": "uuid",
    "roleName": "TEACHER",
    "createdAt": "02-01-2006 15:04:05"
  }
]
```

### Create Role
- **URL:** `/roles`
- **Method:** `POST`
- **Body:**
```json
{
  "roleName": "TEACHER"
}
```

### Get Role by ID
- **URL:** `/roles/:id`
- **Method:** `GET`

### Update Role Name
- **URL:** `/roles/:id`
- **Method:** `PATCH`
- **Body:** `{"roleName": "Wali Kelas"}`

### Delete Role
- **URL:** `/roles/:id`
- **Method:** `DELETE`

---

## 2. User Role Management (Assignments)
Assigning global roles to specific users within a school context.

### Assign Role to User
- **URL:** `/user-roles`
- **Method:** `POST`
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

### List User's Roles
- **URL:** `/user-roles/:schoolUserId`
- **Method:** `GET`

### Update User Roles (Sync)
Replace all roles for a user.
- **URL:** `/user-roles/:schoolUserId`
- **Method:** `PATCH`
- **Body:** `{"roleIds": ["role-uuid-1", "role-uuid-2"]}`
