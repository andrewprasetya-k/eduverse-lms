# 📜 System Logs (Audit Trail) Module API Documentation

Base URL: `/api/logs`

## 1. List Logs by School
Retrieve a paginated history of system activities for a specific school.

- **URL:** `/school/:schoolId`
- **Method:** `GET`
- **Query Parameters:** `page`, `limit`.

**Response Example:**
```json
{
  "data": [
    {
      "logId": "uuid",
      "userId": "uuid",
      "userName": "Admin Budi",
      "action": "CREATE_CLASS",
      "metadata": "{"classTitle": "X-IPA-1"}",
      "createdAt": "13-02-2026 17:00:00"
    }
  ],
  "totalItems": 500,
  ...
}
```
