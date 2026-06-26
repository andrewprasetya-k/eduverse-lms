# 👤 User Profile Module API Documentation

Base URL: `/api/users`

## 1. List All Users
Retrieve a paginated list of all global users.

- **URL:** `(base URL)`
- **Method:** `GET`
- **Auth:** Required (`admin`, `super_admin`)
- **Scope:** Used by school admins to search existing global users before adding them
  as school members. It does not create or mutate global users.
- **Query Parameters:**
  - `page` (default: `1`)
  - `limit` (default: `10`)
  - `search` (optional): Search by full name or email.

**Response Example:**
```json
{
  "data": [...],
  "totalItems": 100,
  "page": 1,
  "limit": 10,
  "totalPages": 10
}
```

---

## 2. Create User
Register a new global user profile. Password will be securely hashed.

- **URL:** `(base URL)`
- **Method:** `POST`
- **Auth:** Required (`super_admin` on system school `sch_code = "000000"` only)
- **Scope:** Global platform account creation. School admins must not create global users
  directly; they can add existing global users to their school through school membership
  and assign school roles through RBAC.
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `fullName` | string | Yes | |
| `email` | string | Yes | Unique |
| `password` | string | Yes | Min 6 characters |

---

## 3. Get User Detail
- **URL:** `/:id`
- **Method:** `GET`
- **Auth:** Required (`super_admin` on system school `sch_code = "000000"` only)

---

## 4. Update User Detail
- **URL:** `/:id`
- **Method:** `PATCH`
- **Auth:** Required (`super_admin` on system school `sch_code = "000000"` only)
- **Body:** `fullName`, `email`.

---

## 5. Change Password
Update user password. Old password verification is required.

- **URL:** `/:id/change-password`
- **Method:** `PATCH`
- **Auth:** Required (`super_admin` on system school `sch_code = "000000"` only)
- **Note:** Current-user password changes should move to a future `/me/change-password`
  endpoint that uses JWT identity instead of path `:id`.
- **Body:**
  - `oldPassword` (string, required)
  - `newPassword` (string, required, min 6)

---

## 6. Delete User
- **URL:** `/:id`
- **Method:** `DELETE`
- **Auth:** Required (`super_admin` on system school `sch_code = "000000"` only)
