# 📝 Assignment & Grading Module API Documentation

Base URL: `/api/assignments`

## Categories

### 1. Create Category
- **URL:** `/categories`
- **Method:** `POST`
- **Body:** `{"schoolId": "uuid", "categoryName": "Kuis"}`

### 2. Get Categories by School
- **URL:** `/categories/school/:schoolCode`
- **Method:** `GET`
- **Response:** `SchoolWithAssignmentCategoriesDTO`

---

## Assignments

### 3. Create Assignment
- **URL:** `/`
- **Method:** `POST`
- **Body:**
```json
{
  "schoolId": "uuid",
  "subjectClassId": "uuid",
  "categoryId": "uuid",
  "assignmentTitle": "string",
  "assignmentDescription": "string",
  "deadline": "2026-03-01T23:59:59Z",
  "allowLateSubmission": false,
  "createdBy": "uuid",
  "mediaIds": ["uuid"]
}
```

### 4. List Assignments by Subject Class
- **URL:** `/subject-class/:subjectClassId`
- **Method:** `GET`
- **Response:** `AssignmentPerSubjectClassResponseDTO` (with subject class header)

### 5. Get Assignment with Submissions
- **URL:** `/:id`
- **Method:** `GET`
- **Response:** `AssignmentWithSubmissionsDTO` (includes all submissions and assessments)

### 6. Update Assignment
- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:** (all fields optional)
```json
{
  "categoryId": "uuid",
  "assignmentTitle": "string",
  "assignmentDescription": "string",
  "deadline": "2026-03-01T23:59:59Z",
  "allowLateSubmission": true,
  "mediaIds": ["uuid"]
}
```

### 7. Delete Assignment
- **URL:** `/:id`
- **Method:** `DELETE`
- **Note:** Soft delete

---

## Submissions

### 8. Submit Assignment
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
- **Note:** Upsert logic - updates existing submission if already submitted

### 9. Get Submission by ID
- **URL:** `/submit/:submissionId`
- **Method:** `GET`
- **Response:** Includes `isLate` indicator and assessment if graded

### 10. Update Submission
- **URL:** `/submit/:submissionId`
- **Method:** `PATCH`
- **Body:**
```json
{
  "schoolId": "uuid",
  "userId": "uuid",
  "mediaIds": ["uuid"]
}
```

### 11. Delete Submission
- **URL:** `/submit/:submissionId`
- **Method:** `DELETE`
- **Note:** Soft delete, can be restored by resubmitting

---

## Assessments

### 12. Grade Submission
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
- **Note:** Upsert logic - updates existing assessment if already graded

### 13. Update Assessment
- **URL:** `/assess/:submissionId`
- **Method:** `PATCH`
- **Body:** (all fields optional)
```json
{
  "score": 95.0,
  "feedback": "Excellent work"
}
```

### 14. Delete Assessment
- **URL:** `/assess/:submissionId`
- **Method:** `DELETE`
- **Note:** Removes grading, submission remains

---

## Key Features

- **Late Submission Control:** `allowLateSubmission` flag per assignment
- **Upsert Logic:** Submissions and assessments auto-update if already exist
- **Soft Delete:** Assignments and submissions can be restored
- **IsLate Indicator:** Automatically calculated in submission responses
- **Attachments:** Support for multiple media files per assignment/submission
