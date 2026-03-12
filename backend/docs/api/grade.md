# 📊 Grade Book API Documentation

Base URL: `/api/grades`

## Overview

Grade Book system untuk manage assessment weights dan calculate final grades berdasarkan weighted categories.

---

## 1. Configure Assessment Weights
Set bobot penilaian per kategori untuk mata pelajaran.

- **URL:** `/weights`
- **Method:** `POST`
- **Auth:** Required (admin, teacher)
- **Body:**
```json
{
  "subjectId": "uuid-subject",
  "weights": [
    {
      "categoryId": "uuid-quiz",
      "weight": 20.00
    },
    {
      "categoryId": "uuid-uts", 
      "weight": 30.00
    },
    {
      "categoryId": "uuid-uas",
      "weight": 50.00
    }
  ]
}
```

**Validation:**
- Total weight harus = 100.00
- Weight per kategori: 0-100
- Minimum 1 kategori

**Response (200 OK):**
```json
{
  "message": "Weights configured successfully"
}
```

---

## 2. Get Assessment Weights by Subject
Retrieve konfigurasi bobot untuk mata pelajaran.

- **URL:** `/weights/subject/:subjectId`
- **Method:** `GET`
- **Auth:** Required (school member)

**Response (200 OK):**
```json
{
  "subjectId": "uuid",
  "subjectName": "Matematika",
  "subjectCode": "MTK",
  "weights": [
    {
      "weightId": "uuid",
      "categoryId": "uuid",
      "categoryName": "Quiz",
      "weight": 20.00
    },
    {
      "weightId": "uuid", 
      "categoryId": "uuid",
      "categoryName": "UTS",
      "weight": 30.00
    }
  ],
  "totalWeight": 100.00
}
```

---

## 3. Get Student Final Grade
Calculate dan retrieve final grade untuk student di mata pelajaran tertentu.

- **URL:** `/student/:userId/subject/:subjectId`
- **Method:** `GET`
- **Auth:** Required (school member)

**Response (200 OK):**
```json
{
  "studentId": "uuid",
  "studentName": "John Doe",
  "subjectId": "uuid",
  "subjectName": "Matematika",
  "breakdown": [
    {
      "categoryId": "uuid",
      "categoryName": "Quiz",
      "weight": 20.00,
      "averageScore": 85.00,
      "weightedScore": 17.00,
      "assignmentCount": 3
    },
    {
      "categoryId": "uuid",
      "categoryName": "UTS", 
      "weight": 30.00,
      "averageScore": 90.00,
      "weightedScore": 27.00,
      "assignmentCount": 1
    }
  ],
  "finalGrade": 82.50,
  "letterGrade": "A"
}
```

---

## 4. Get Class Grade Report
Retrieve final grades untuk seluruh student di kelas untuk mata pelajaran tertentu.

- **URL:** `/class/:classId/subject/:subjectId`
- **Method:** `GET`
- **Auth:** Required (teacher, admin)

**Response (200 OK):**
```json
{
  "class": {
    "classId": "uuid",
    "className": "12 IPA 1",
    "classCode": "12IPA1"
  },
  "subject": {
    "subjectId": "uuid",
    "subjectName": "Matematika",
    "subjectCode": "MTK"
  },
  "students": [
    {
      "studentId": "uuid",
      "studentName": "John Doe",
      "studentEmail": "john@example.com",
      "finalGrade": 82.50,
      "letterGrade": "A"
    },
    {
      "studentId": "uuid",
      "studentName": "Jane Smith", 
      "studentEmail": "jane@example.com",
      "finalGrade": 78.25,
      "letterGrade": "B"
    }
  ]
}
```

---

## Letter Grade Conversion

| Score Range | Letter Grade |
|-------------|--------------|
| 90-100      | A            |
| 80-89       | B            |
| 70-79       | C            |
| 60-69       | D            |
| 0-59        | E            |

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Total weight must be 100, got 95.00"
}
```

### 404 Not Found
```json
{
  "error": "No weights configured for this subject"
}
```

---

## Usage Flow

1. **Admin/Teacher configure weights** per subject
2. **Teacher grade assignments** (via assignment endpoints)
3. **System auto-calculate final grades** based on weights
4. **Students/Teachers view final grades** with breakdown

---

**Last Updated:** 2026-03-12
