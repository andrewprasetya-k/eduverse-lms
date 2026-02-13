# 🏫 School User (Membership) Module API Documentation

Base URL: `/api/school-users`

## 1. Enroll User to School
Connect an existing user to a school.

- **URL:** `/enroll`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `userId` | uuid | Yes | |
| `schoolId` | uuid | Yes | |

---

## 2. List Members by School
Retrieve all users who belong to a specific school.

- **URL:** `/school/:schoolId`
- **Method:** `GET`

---

## 3. List Schools by User
Retrieve all schools that a specific user belongs to.

- **URL:** `/user/:userId`
- **Method:** `GET`

---

## 4. Unenroll User
Remove a user's membership from a school.

- **URL:** `/:id`
- **Method:** `DELETE`
