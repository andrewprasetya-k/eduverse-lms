# 📁 Media & Metadata Module API Documentation

Base URL: `/api/medias`

## 1. Record Media Metadata
After uploading a file to external storage (like S3 or Supabase), use this endpoint to record the metadata in the database.

- **URL:** `/metadata`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `mediaName`| string | Yes | |
| `fileSize` | int64 | Yes | In bytes |
| `mimeType` | string | Yes | e.g., "image/png" |
| `storagePath`| string | Yes | Path in S3/Supabase |
| `fileUrl` | string | Yes | Public URL |
| `thumbnailUrl`| string | No | |
| `ownerType`| string | Yes | `user`, `school`, `material`, etc. |
| `ownerId` | uuid | Yes | |

---

## 2. Get Media Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 3. Delete Media Record
- **URL:** `/:id`
- **Method:** `DELETE`
