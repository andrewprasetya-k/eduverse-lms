# 🏫 School Module API Documentation

Base URL: `/api/schools`

## 1. List Schools
Retrieve a paginated list of schools with filtering, searching, and sorting capabilities.

- **URL:** `/`
- **Method:** `GET`
- **Query Parameters:**
  - `page` (default: `1`): Page number.
  - `limit` (default: `10`): Items per page.
  - `search` (optional): Filter by school name or code.
  - `status` (optional): `active` (default), `deleted` (in trash), `all`.
  - `sortBy` (optional): Sort column using these keys:
    - `name` (maps to school name)
    - `code` (maps to school code)
    - `createdAt` (maps to creation date, default)
    - `updatedAt` (maps to last update date)
  - `order` (optional): `asc` (A-Z) or `desc` (Z-A, newest, default).

**Example Request:**
`GET /api/schools?search=edu&status=active&sortBy=name&order=asc`

---

## 2. Get School Summary
Get high-level statistics for school management.

- **URL:** `/summary`
- **Method:** `GET`

**Response Example:**
```json
{
  "totalActive": 10,
  "totalDeleted": 2,
  "totalSchools": 12
}
```

---

## 3. Check Code Availability
Quickly check if a school code is already taken before submitting a form.

- **URL:** `/check-code/:schoolCode`
- **Method:** `GET`

**Response Example:**
```json
{
  "schoolCode": "EDU01",
  "available": true
}
```

---

## 4. Create School
Register a new school in the system.

- **URL:** `/`
- **Method:** `POST`
- **Payload:**
```json
{
  "schoolName": "Eduverse Academy",
  "schoolCode": "EDU01",
  "schoolAddress": "Jl. Merdeka No. 1",
  "schoolEmail": "admin@edu.com",
  "schoolPhone": "081234567890",
  "schoolWebsite": "https://edu.com"
}
```

---

## 5. Get School Detail
Get full information of a specific school by its code.

- **URL:** `/:schoolCode`
- **Method:** `GET`

---

## 6. Update School
Update existing school information.

- **URL:** `/:schoolCode`
- **Method:** `PATCH`
- **Payload:** (Partial updates are allowed)

---

## 7. Management Actions
- **Soft Delete:** `DELETE /:schoolCode` (Moves to trash)
- **Restore:** `PATCH /restore/:schoolCode` (Brings back from trash)
- **Hard Delete:** `DELETE /:schoolCode/permanent` (Permanently removes from DB)
