# Chat API

Base URL: `/api/chat`

REST Chat MVP hanya mendukung chat per `subject_class`.

## Scope MVP

- Subject-class chat only.
- Text-only messages.
- REST API only.
- Student dan teacher only.
- Tidak ada WebSocket/realtime.
- Tidak ada DM, free group chat, attachment, typing indicator, online/offline,
  delete/unsend, reply UI, admin moderation, atau notification integration.

## Access Rules

Semua endpoint memerlukan:

- JWT authentication.
- Active `SchoolId` context.
- Active school membership (`school_users.deleted_at IS NULL`).
- Role `student` atau `teacher`.

Room permission tidak memakai `chat_room_members` sebagai source of truth untuk
MVP. Akses diturunkan dari aturan akademik:

- Student boleh mengakses room subject-class jika aktif terdaftar di class
  subject-class tersebut (`enrollments.left_at IS NULL`).
- Teacher boleh mengakses room subject-class jika aktif mengajar subject-class
  tersebut.
- `chat_rooms.room_sch_id` harus sama dengan active school.
- `chat_rooms.deleted_at IS NULL`.
- Admin dan super admin tidak otomatis ikut chat akademik MVP.

## Endpoints

### List My Rooms

`GET /rooms`

Mengembalikan room subject-class yang sudah dibuka dan dapat diakses oleh user
saat ini.

```json
{
  "rooms": [
    {
      "roomId": "uuid",
      "subjectClassId": "uuid",
      "subjectId": "uuid",
      "subjectName": "Matematika",
      "subjectCode": "MTK",
      "classId": "uuid",
      "className": "Kelas 10 IPA",
      "classCode": "10-IPA",
      "roomName": "Matematika - Kelas 10 IPA",
      "lastMessage": {
        "messageId": "uuid",
        "senderId": "uuid",
        "senderName": "Budi",
        "content": "Baik, Bu.",
        "createdAt": "2026-06-26T03:00:00Z"
      },
      "lastMessageAt": "2026-06-26T03:00:00Z",
      "unreadCount": 1,
      "canSend": true
    }
  ]
}
```

### Open Subject-Class Room

`POST /subject-classes/:subjectClassId/open`

Membuka atau membuat room subject-class. Satu subject-class hanya memiliki satu
room per school.

```json
{
  "room": {
    "roomId": "uuid",
    "subjectClassId": "uuid",
    "subjectName": "Matematika",
    "className": "Kelas 10 IPA",
    "roomName": "Matematika - Kelas 10 IPA",
    "unreadCount": 0,
    "canSend": true
  }
}
```

### List Messages

`GET /rooms/:roomId/messages?limit=50&before=2026-06-26T03:00:00Z`

`limit` dibatasi maksimal 50. Response diurutkan oldest-to-newest untuk
keterbacaan percakapan. `nextBefore` bisa dipakai untuk mengambil pesan yang
lebih lama.

```json
{
  "messages": [
    {
      "messageId": "uuid",
      "roomId": "uuid",
      "senderId": "uuid",
      "senderName": "Budi",
      "senderRole": "student",
      "content": "Baik, Bu.",
      "messageType": "text",
      "createdAt": "2026-06-26T03:00:00Z",
      "isMine": true
    }
  ],
  "nextBefore": null,
  "hasMore": false
}
```

### Create Message

`POST /rooms/:roomId/messages`

```json
{
  "content": "Halo kelas."
}
```

Rules:

- Content di-trim.
- Empty content ditolak.
- Maksimal 5.000 karakter.
- `messageType` selalu `text`.

Response adalah canonical `MessageDTO` dan dapat dipakai ulang nanti sebagai
payload WebSocket `new_message`.

### Mark Room Read

`PATCH /rooms/:roomId/read`

```json
{
  "lastReadMessageId": "uuid"
}
```

`lastReadMessageId` opsional. Endpoint ini idempotent dan hanya berlaku jika
current user memiliki akses ke room.

