# 🔐 RBAC (Roles & Permissions) Module API Documentation

Base URL: `/api/rbac`

## 1. Role Management

### List Roles by School
- **URL:** `/roles/school/:schoolCode`
- **Method:** `GET`

**Response Example:**
```json
[
  {
    "roleId": "uuid",
    "school": {
      "schoolId": "uuid",
      "schoolName": "Eduverse Academy",
      "schoolCode": "SD-ANGKASA"
    },
    "roleName": "Guru",
    "createdAt": "..."
  }
]
```

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
- **URL:** `/roles/permissions/:id`
- **Method:** `PATCH`
- **Body:** `{"permissionIds": ["uuid-1", "uuid-3"]}`

---

## 2. Permission Management
Retrieve or create global permission keys.

### Create Permission
Create a new global permission key.
- **URL:** `/permissions`
- **Method:** `POST`
- **Body:**
```json
{
  "permissionKey": "COURSE_CREATE",
  "description": "Ability to create new courses"
}
```

### List All Available Permissions
Retrieve global permission keys.
- **URL:** `/permissions`
- **Method:** `GET`

### Get Permission by ID
- **URL:** `/permissions/:id`
- **Method:** `GET`

### Update Permission
- **URL:** `/permissions/:id`
- **Method:** `PATCH`
- **Body:**
```json
{
  "permissionKey": "COURSE_EDIT",
  "permissionDesc": "Updated description"
}
```

### Delete Permission
- **URL:** `/permissions/:id`
- **Method:** `DELETE`

---

## 3. User Role Management (Assignments)

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
- **URL:** `/user-roles/user/:schoolUserId`
- **Method:** `GET`

### Update User Roles (Sync)
Replace all roles for a user.
- **URL:** `/user-roles/user/:schoolUserId`
- **Method:** `PATCH`
- **Body:** `{"roleIds": ["role-uuid-1", "role-uuid-2"]}`
