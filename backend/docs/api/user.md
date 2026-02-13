# 👤 User Profile Module API Documentation

Base URL: `/api/users`

## 1. Create User
Register a new global user profile. Password will be securely hashed.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `fullName` | string | Yes | |
| `email` | string | Yes | Unique |
| `password` | string | Yes | Min 6 characters |

---

## 2. Get User Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 3. Update User Detail
- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:** `fullName`, `email`.

---

## 4. Delete User
- **URL:** `/:id`
- **Method:** `DELETE`
