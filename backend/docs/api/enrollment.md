# 🎓 Enrollment (Pendaftaran Siswa) Module API Documentation

Base URL: `/api/enrollments`

## 1. Enroll User to Class
Register a school member (student or teacher) into a specific class.

- **URL:** `/enroll`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `schoolUserId`| uuid | Yes | Reference to school_users |
| `classId` | uuid | Yes | |
| `role` | string | Yes | `teacher` or `student` |

---

## 2. List Participants in Class
Retrieve all members enrolled in a specific class.

- **URL:** `/class/:classId`
- **Method:** `GET`

**Response Example:**
```json
[
  {
    "enrollmentId": "uuid",
    "schoolUserId": "uuid",
    "userFullName": "John Doe",
    "userEmail": "john@doe.com",
    "role": "student",
    "joinedAt": "13-02-2026 16:00:00"
  }
]
```

---

## 3. List Classes by Member
Retrieve all classes a specific school member is enrolled in.

- **URL:** `/member/:schoolUserId`
- **Method:** `GET`

---

## 4. Remove Enrollment (Unenroll)
- **URL:** `/:id`
- **Method:** `DELETE`
