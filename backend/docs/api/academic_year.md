# 📅 Academic Year Module API Documentation

Base URL: `/api/academic-years`

## 1. List Academic Years by School
Retrieve all academic years for a specific school using its code.

- **URL:** `/school/:schoolCode`
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

## 2. Get Academic Year Detail
Retrieve detail of a specific academic year by its ID.

- **URL:** `/:id`
- **Method:** `GET`

---

## 3. Create Academic Year
Create a new academic year for a school.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | Reference to School ID |
| `academicYearName` | string | Yes | e.g., "2023/2024" |
| `isActive` | boolean| No | Automatically set to `true` on creation |

---

## 4. Update Academic Year
Update name or activation status.

- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:** (Partial updates allowed)
  - `academicYearName` (string)
  - `isActive` (boolean) - If set to `true`, other years in the same school will be deactivated.

---

## 5. Delete Academic Year
Permanently remove an academic year.

- **URL:** `/:id`
- **Method:** `DELETE`
