# 📣 Classroom Feed Module API Documentation

Base URL: `/api/feeds`

## 1. Post to Feed
Post an announcement or discussion to a class.

- **URL:** `/`
- **Method:** `POST`
- **Body:**
| Field | Type | Required | Note |
| :--- | :--- | :--- | :--- |
| `schoolId` | uuid | Yes | |
| `classId` | uuid | Yes | |
| `content` | string | Yes | Markdown supported |
| `createdBy` | uuid | Yes | |
| `mediaIds` | uuid[] | No | Attachments |

---

## 2. List Class Feeds
- **URL:** `/class/:classId`
- **Method:** `GET`
- **Query Params:** `page`, `limit`.
