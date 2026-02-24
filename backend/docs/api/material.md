# 📖 Material (Konten Belajar) Module API Documentation

Base URL: `/api/materials`

## 1. Create Material
Create a new learning material for a class with optional attachments.

- **URL:** `/`
- **Method:** `POST`

### Option A: JSON (with existing media IDs or inline media data)
- **Content-Type:** `application/json`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `subjectClassId` | uuid | Yes | Link to subject, class, and teacher |
| `materialTitle`| string | Yes | |
| `materialDesc` | string | No | |
| `materialType`| string | Yes | `video`, `pdf`, `ppt`, `other` |
| `createdBy` | uuid | Yes | ID of User (Teacher/Admin) |
| `mediaIds` | uuid[] | No | List of already recorded Media IDs |
| `medias` | object[] | No | Inline media data (auto-create in medias table) |

**Inline Media Object:**
```json
{
  "name": "filename.pdf",
  "fileSize": 1024000,
  "mimeType": "application/pdf",
  "fileUrl": "https://supabase.co/storage/.../file.pdf",
  "thumbnailUrl": "https://..." // optional
}
```

### Option B: Multipart Form (with file uploads)
- **Content-Type:** `multipart/form-data`
- **Form Fields:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | string | Yes | UUID |
| `subjectClassId` | string | Yes | UUID |
| `materialTitle`| string | Yes | |
| `materialDesc` | string | No | |
| `materialType`| string | Yes | `video`, `pdf`, `ppt`, `other` |
| `createdBy` | string | Yes | UUID |
| `files` | file[] | No | Multiple files (auto-detect size, mime type) |

---

## 2. List Materials
- **URL:** `/`
- **Method:** `GET`
- **Query Params:** `page`, `limit`, `search`, `subjectClassId`.
- **Note:** If `subjectClassId` is provided, response will be wrapped in `MaterialListWithSubjectDTO`.

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
| `materialDesc` | string | Optional |
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
