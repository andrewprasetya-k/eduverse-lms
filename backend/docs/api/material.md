# 📖 Material (Konten Belajar) Module API Documentation

Base URL: `/api/materials`

## 1. Create Material
Create a new learning material for a class with optional attachments.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `classId` | uuid | Yes | |
| `materialTitle`| string | Yes | |
| `materialDescription` | string | No | |
| `materialType`| string | Yes | `video`, `pdf`, `ppt`, `other` |
| `createdBy` | uuid | Yes | ID of User (Teacher/Admin) |
| `mediaIds` | uuid[] | No | List of already recorded Media IDs |

---

## 2. List Materials
- **URL:** `/`
- **Method:** `GET`
- **Query Params:** `page`, `limit`, `search`, `classId`.

---

## 3. Get Material Detail (with Attachments)
- **URL:** `/:id`
- **Method:** `GET`

---

## 4. Update Material
Update material details and its attachments.

- **URL:** `/:id`
- **Method:** `PATCH`
- **Body:**
| Field | Type | Note |
| :--- | :--- | :--- |
| `materialTitle`| string | Optional |
| `materialDescription` | string | Optional |
| `materialType`| string | `video`, `pdf`, `ppt`, `other` (Optional) |
| `mediaIds` | uuid[] | New list of Media IDs (Will replace existing) |

---

## 5. Delete Material
- **URL:** `/:id`
- **Method:** `DELETE`

---

## 6. Update Progress
Mark a material as completed for a user.

- **URL:** `/progress`
- **Method:** `POST`
- **Body:**
```json
{
  "userId": "uuid",
  "materialId": "uuid",
  "status": "completed"
}
```
