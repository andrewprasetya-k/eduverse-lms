# 💬 Comments Module API Documentation

Base URL: `/api/comments`

## 1. Post a Comment
- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `sourceType`| string | Yes | `feed`, `assignment`, etc. |
| `sourceId` | uuid | Yes | |
| `userId` | uuid | Yes | |
| `content` | string | Yes | |

---

## 2. List Comments by Source
- **URL:** `/`
- **Method:** `GET`
- **Query Params:**
  - `type`: e.g., `feed`
  - `id`: The source UUID.

---

## 3. Delete Comment
- **URL:** `/:id`
- **Method:** `DELETE`
