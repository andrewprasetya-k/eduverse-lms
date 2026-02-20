# 🎓 Enrollment (Pendaftaran Siswa) Module API Documentation

Base URL: `/api/enrollments`

## 1. Enroll Users to Class
Register one or more school members (students or teachers) into a specific class.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `schoolUserIds`| uuid[] | Yes | List of references to school_users |
| `classId` | uuid | Yes | |
| `role` | string | Yes | `teacher` or `student` |

---

## 2. List Participants in Class
Retrieve all members enrolled in a specific class.

- **URL:** `/class/:classId`
- **Method:** `GET`
- **Response:** `ClassWithMembersDTO` (Includes class header and list of members)

---

## 3. Get Enrollment Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 4. List Classes by Member
Retrieve all classes a specific school member is enrolled in.

- **URL:** `/member/:schoolUserId`
- **Method:** `GET`

---

## 5. Remove Enrollment (Unenroll)
- **URL:** `/:id`
- **Method:** `DELETE`
