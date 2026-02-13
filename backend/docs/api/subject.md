# 📚 Subject (Mata Pelajaran) Module API Documentation

Base URL: `/api/subjects`

## 1. List All Subjects
Retrieve a paginated list of all subjects (Super Admin view).

- **URL:** `/`
- **Method:** `GET`
- **Query Parameters:**
  - `page` (default: `1`)
  - `limit` (default: `10`)
  - `search` (optional): Search by name or code.

---

## 2. List Subjects by School
Retrieve all subjects for a specific school.

- **URL:** `/school/:schoolCode`
- **Method:** `GET`

---

## 3. Get Subject Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 4. Create Subject
Register a new subject for a school.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `subjectName`| string | Yes | e.g., "Matematika" |
| `subjectCode`| string | Yes | Unique per school, e.g., "MTK" |

---

## 5. Update Subject
- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:** `subjectName`, `subjectCode`.

---

## 6. Delete Subject
- **URL:** `/:id`
- **Method:** `DELETE`
