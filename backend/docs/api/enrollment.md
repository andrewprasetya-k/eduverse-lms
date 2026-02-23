# 👥 Enrollment Module API Documentation

Base URL: `/api/enrollments`

## 1. Enroll Members to Class
- **URL:** `/`
- **Method:** `POST`
- **Body:**
```json
{
  "schoolId": "uuid",
  "schoolUserIds": ["uuid1", "uuid2"],
  "classId": "uuid",
  "role": "teacher|student"
}
```
- **Note:** Bulk enrollment supported

## 2. Get Enrollments by Class
- **URL:** `/class/:classId`
- **Method:** `GET`
- **Response:** `ClassWithMembersDTO` (with class header and member list)

## 3. Get Enrollments by Member
- **URL:** `/member/:schoolUserId`
- **Method:** `GET`
- **Response:** List of classes the member is enrolled in

## 4. Get Enrollment by ID
- **URL:** `/:id`
- **Method:** `GET`
- **Response:** Single enrollment with user and class details

## 5. Update Enrollment Role
- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:**
```json
{
  "role": "teacher|student"
}
```
- **Use Case:** Change member role (e.g., promote student to teacher assistant)

## 6. Unenroll Member
- **URL:** `/:id`
- **Method:** `DELETE`
- **Note:** Removes member from class

---

## Features

- **Bulk Enrollment:** Multiple users can be enrolled at once
- **Role Management:** Support for teacher and student roles
- **Bidirectional Queries:** Get by class or by member
- **Class Context:** Enrollment list includes class header
