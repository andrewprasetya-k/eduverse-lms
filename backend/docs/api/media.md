# 📁 Media & Metadata Module API Documentation

Base URL: `/api/medias`

## 1. Upload File
Upload file directly to backend (multipart form). System will auto-detect file size, mime type, and filename.

- **URL:** `/upload`
- **Method:** `POST`
- **Content-Type:** `multipart/form-data`
- **Form Fields:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `file` | file | Yes | The file to upload |
| `schoolId` | string | Yes | UUID |
| `ownerType`| string | No | `user`, `school`, `material`, etc. |
| `ownerId` | string | No | UUID |

**Response:**
```json
{
  "message": "File uploaded successfully",
  "mediaId": "uuid",
  "fileName": "example.pdf",
  "fileSize": 1024000,
  "mimeType": "application/pdf",
  "fileUrl": "https://...",
  "ext": ".pdf"
}
```

**Note:** Currently returns placeholder URL. For production, integrate with Supabase Storage client.

---

## 2. Record Media Metadata
After uploading a file to external storage (like S3 or Supabase), use this endpoint to record the metadata in the database.

- **URL:** `/metadata`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `name`| string | Yes | |
| `fileSize` | int64 | Yes | In bytes |
| `mimeType` | string | Yes | e.g., "image/png" |
| `storagePath`| string | Yes | Path in S3/Supabase |
| `fileUrl` | string | Yes | Public URL |
| `thumbnailUrl`| string | No | |
| `isPublic` | boolean | No | Default: true |
| `ownerType`| string | Yes | `user`, `school`, `material`, etc. |
| `ownerId` | uuid | Yes | |

---

## 3. Get Media Detail
- **URL:** `/:id`
- **Method:** `GET`

---

## 4. Delete Media Record
- **URL:** `/:id`
- **Method:** `DELETE`
