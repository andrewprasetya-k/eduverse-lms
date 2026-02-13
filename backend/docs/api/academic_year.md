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
Create a new academic year for a school. Status is `false` by default.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | Reference to School ID |
| `academicYearName` | string | Yes | e.g., "2023/2024" |

---

## 4. Update Academic Year
Update basic information of an academic year.

- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:**
  - `academicYearName` (string)

---

## 5. Activate Academic Year
Set an academic year as the active one for its school. This will automatically deactivate all other academic years in the same school.

- **URL:** `/activate/:id`
- **Method:** `POST`
- **Response:** `{"message": "Academic year activated successfully"}`

---

## 6. Deactivate Academic Year
Manually deactivate an academic year.

- **URL:** `/deactivate/:id`
- **Method:** `POST`
- **Response:** `{"message": "Academic year deactivated successfully"}`

---

## 7. Delete Academic Year
Permanently remove an academic year.

- **URL:** `/:id`
- **Method:** `DELETE`
