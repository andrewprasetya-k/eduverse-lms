# 📅 Academic Year Module API Documentation

Base URL: `/api/academic-years`

## 1. List Academic Years by School
Retrieve all academic years for a specific school.

- **URL:** `/school/:schoolId`
- **Method:** `GET`

**Response Example:**
```json
[
  {
    "academicYearId": "uuid-string",
    "schoolId": "uuid-school-id",
    "academicYearName": "2023/2024",
    "isActive": true,
    "createdAt": "13-02-2026 10:00:00"
  }
]
```

---

## 2. Create Academic Year
Create a new academic year for a school.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | Reference to School ID |
| `academicYearName` | string | Yes | e.g., "2023/2024" |
| `isActive` | boolean| No | If true, other years in the same school will be deactivated |

---

## 3. Update Academic Year
Update name or activation status.

- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:** (Partial updates allowed)
  - `academicYearName` (string)
  - `isActive` (boolean)

---

## 4. Delete Academic Year
Permanently remove an academic year.

- **URL:** `/:id`
- **Method:** `DELETE`
