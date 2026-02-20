# 📝 Assignment & Grading Module API Documentation

Base URL: `/api/assignments`

## 1. Create Category
- **URL:** `/categories`
- **Method:** `POST`
- **Body:** `{"schoolId": "uuid", "categoryName": "Kuis"}`

---

## 2. Get Categories by School
- **URL:** `/categories/school/:schoolCode`
- **Method:** `GET`
- **Response:** `SchoolWithAssignmentCategoriesDTO`

---

## 3. Create Assignment
- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `subjectClassId` | uuid | Yes | Link to subject, class, and teacher |
| `categoryId` | uuid | Yes | |
| `assignmentTitle`| string | Yes | |
| `assignmentDescription`| string | No | |
| `deadline` | datetime | No | ISO format (e.g. 2026-03-01T23:59:59Z) |
| `createdBy` | uuid | Yes | Teacher/Admin ID |
| `mediaIds` | uuid[] | No | Attachments |

---

## 4. List Subject-Class Assignments
- **URL:** `/subject-class/:subjectClassId`
- **Method:** `GET`

---

## 5. Get Assignment Submissions (Monitoring)
- **URL:** `/:id/submissions`
- **Method:** `GET`
- **Response:** `AssignmentWithSubmissionsDTO` (Includes header and list of submissions with assessments)

---

## 6. Submit Assignment (Upsert)
- **URL:** `/submit/:assignmentId`
- **Method:** `POST`
- **Body:** 
```json
{
  "schoolId": "uuid",
  "userId": "uuid",
  "mediaIds": ["uuid"]
}
```
*Note: If a student submits again, the existing record will be updated (Upsert).*

---

## 7. Record Assessment (Grade)
- **URL:** `/assess/:submissionId`
- **Method:** `POST`
- **Body:**
```json
{
  "score": 90.5,
  "feedback": "Good job",
  "assessedBy": "uuid"
}
```
*Note: Existing assessment will be updated if already graded.*
