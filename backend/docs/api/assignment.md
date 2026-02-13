# 📝 Assignment & Grading Module API Documentation

Base URL: `/api/assignments`

## 1. Create Category
- **URL:** `/categories`
- **Method:** `POST`
- **Body:** `{"schoolId": "uuid", "categoryName": "Kuis"}`

---

## 2. Create Assignment
- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `classId` | uuid | Yes | |
| `categoryId` | uuid | Yes | |
| `assignmentTitle`| string | Yes | |
| `deadline` | datetime | No | ISO format |
| `mediaIds` | uuid[] | No | Attachments |

---

## 3. List Class Assignments
- **URL:** `/class/:classId`
- **Method:** `GET`

---

## 4. Submit Assignment
- **URL:** `/submit`
- **Method:** `POST`
- **Body:** `{"schoolId": "uuid", "assignmentId": "uuid", "userId": "uuid", "mediaIds": ["uuid"]}`

---

## 5. Record Assessment (Grade)
- **URL:** `/assess`
- **Method:** `POST`
- **Body:** `{"submissionId": "uuid", "score": 90.5, "feedback": "Good job", "assessedBy": "uuid"}`
