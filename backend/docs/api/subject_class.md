# 📖 Subject Class (Penugasan Guru) Module API Documentation

Base URL: `/api/subject-classes`

## 1. Assign Subject and Teacher to Class
Link a specific subject and teacher to a class.

- **URL:** `/assign`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `classId` | uuid | Yes | |
| `subjectId` | uuid | Yes | |
| `teacherId` | uuid | Yes | Reference to school_users (ID Guru) |

---

## 2. List Subjects in Class
Retrieve all subjects and their teachers for a specific class.

- **URL:** `/class/:classId`
- **Method:** `GET`
- **Response:** `SubjectPerClassDTO` (Includes class header and list of subject assignments)

---

## 3. Get Assignment Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 4. Update Assignment
Update teacher or subject for an existing assignment.

- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `subjectId` | uuid | No | |
| `teacherId` | uuid | No | |

---

## 5. Remove Assignment (Unassign)
- **URL:** `/:id`
- **Method:** `DELETE`
