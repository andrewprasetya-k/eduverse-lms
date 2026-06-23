# 📰 Feed Module API Documentation

Base URL: `/api/feeds`

## 1. Create Feed
- **URL:** `(base URL)`
- **Method:** `POST`
- **Auth:** Required. Active school member with `teacher` or `admin` role.
- **Ownership:** Active `SchoolId` header is the source of truth. `schoolId` in the body must match the active school.
- **Teacher rule:** Teacher can create a feed only for a class where the teacher teaches at least one active subject_class.
- **Admin rule:** Admin can create a feed for active-school classes.
- **Body:**
```json
{
  "schoolId": "uuid",
  "classId": "uuid",
  "content": "Announcement text"
}
```

`mediaIds` / feed attachments are not supported in the current MVP.

## 2. Get Feeds by Class
- **URL:** `/class/:classId`
- **Method:** `GET`
- **Auth:** Required. Active school member with `admin`, `teacher`, or `student` role.
- **Access:**
  - Admin can read feeds for active-school classes.
  - Teacher can read feeds for classes they actively teach.
  - Student can read feeds only for classes where they have active enrollment (`left_at IS NULL`).
- **Query Params:** `?page=1&limit=10`
- **Response:** `ClassWithFeedsDTO` (with class header and paginated feeds)

## 3. Get Feed by ID
- **URL:** `/:id`
- **Method:** `GET`
- **Auth:** Required. Same active-school/class access rules as class feed list.
- **Response:** Includes attachments and comment count

## 4. Update Feed
- **URL:** `/:id`
- **Method:** `PATCH`
- **Auth:** Required. Active school member with `teacher` or `admin` role.
- **Access:**
  - Admin can update active-school feed posts.
  - Teacher can update only their own feed posts in classes they actively teach.
- **Body:** (all fields optional)
```json
{
  "content": "Updated announcement"
}
```

`mediaIds` / feed attachments are not supported in the current MVP.

## 5. Delete Feed
- **URL:** `/:id`
- **Method:** `DELETE`
- **Auth:** Required. Active school member with `teacher` or `admin` role.
- **Access:**
  - Admin can delete active-school feed posts.
  - Teacher can delete only their own feed posts in classes they actively teach.
- **Note:** Soft delete

---

## Features

- **Pagination:** Supports page and limit query params
- **Comment Count:** Automatically included in response
- **Class Context:** Feed list includes class header
- **Notifications:** `feed_posted` notification remains best-effort and does not block feed creation.
- **Deferred:** Comments UI, reactions, realtime/WebSocket, and feed attachments are outside the current MVP.
