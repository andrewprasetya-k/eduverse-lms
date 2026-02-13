# 🔐 RBAC (Roles & Permissions) Module API Documentation

Base URL: `/api/rbac`

## 1. Role Management

### List Roles by School
- **URL:** `/roles/school/:schoolCode`
- **Method:** `GET`

### Create Role
- **URL:** `/roles`
- **Method:** `POST`
- **Body:**
```json
{
  "schoolId": "uuid",
  "roleName": "Guru",
  "permissionIds": ["uuid-1", "uuid-2"] // Optional
}
```

### Update Role Name
- **URL:** `/roles/:id`
- **Method:** `PATCH`
- **Body:** `{"roleName": "Wali Kelas"}`

### Set Role Permissions (Sync)
Replace all permissions for a role.
- **URL:** `/roles/:id/permissions`
- **Method:** `PATCH`
- **Body:** `{"permissionIds": ["uuid-1", "uuid-3"]}`

---

## 2. Permission Management (Read Only)

### List All Available Permissions
Retrieve global permission keys.
- **URL:** `/permissions`
- **Method:** `GET`

---

## 3. User Role Assignment

### Assign Role to User
- **URL:** `/assignments`
- **Method:** `POST`
- **Body:**
```json
{
  "schoolUserId": "uuid",
  "roleId": "uuid"
}
```

### Remove Role from User
- **URL:** `/assignments?schoolUserId=...&roleId=...`
- **Method:** `DELETE`

### List User's Roles
- **URL:** `/assignments/user/:schoolUserId`
- **Method:** `GET`
