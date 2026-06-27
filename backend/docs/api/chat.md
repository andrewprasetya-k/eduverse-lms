# Chat API

Base URL: `/api/chat`

REST Chat MVP mendukung satu chat room utama per sekolah aktif dan custom group
room untuk warga aktif di sekolah yang sama.

## Scope MVP

- School-wide chat selalu tersedia sebagai room utama sekolah.
- Custom group room dapat dibuat oleh warga aktif sekolah.
- Text-only messages.
- REST API only.
- Admin Sekolah, teacher, dan student boleh berpartisipasi jika masih menjadi
  member aktif sekolah tersebut.
- Tidak ada WebSocket/realtime.
- Tidak ada subject/class room, DM, attachment, typing indicator,
  online/offline, delete/unsend, moderation UI, atau notification integration.

Subject/class chat adalah ekspansi masa depan.

## Access Rules

Semua endpoint memerlukan:

- JWT authentication.
- Active `SchoolId` context.
- Active school membership (`school_users.deleted_at IS NULL`).

School room permission berbasis sekolah aktif:

- `chat_rooms.room_sch_id` harus sama dengan active school.
- `chat_rooms.room_type = "group"`.
- `chat_rooms.room_ref_type = "school"`.
- `chat_rooms.room_ref_id = activeSchoolID`.
- `chat_rooms.deleted_at IS NULL`.
- Super admin tidak ikut chat akademik sekolah kecuali juga memiliki membership
  aktif di sekolah tersebut.

Custom group room permission:

- `chat_rooms.room_sch_id` harus sama dengan active school.
- `chat_rooms.room_type = "group"`.
- `chat_rooms.room_ref_type IS NULL`.
- `chat_rooms.room_ref_id IS NULL`.
- User harus active school member.
- User juga harus active `chat_room_members` dengan `left_at IS NULL`.

## Endpoints

### List My Rooms

`GET /rooms`

Mengembalikan room sekolah dan custom group room yang dapat diakses oleh user
saat ini.

```json
{
  "rooms": [
    {
      "roomId": "uuid",
      "roomName": "Ruang sekolah",
      "roomType": "group",
      "roomRefType": "school",
      "roomRefId": "school-uuid",
      "schoolId": "school-uuid",
      "schoolName": "SMA EduVerse",
      "lastMessage": {
        "messageId": "uuid",
        "senderId": "uuid",
        "senderName": "Budi",
        "content": "Selamat pagi.",
        "createdAt": "2026-06-26T03:00:00Z"
      },
      "lastMessageAt": "2026-06-26T03:00:00Z",
      "unreadCount": 1,
      "canSend": true
    }
  ]
}
```

### List Chat Members

`GET /members?search=nama`

Mengembalikan warga aktif di active school untuk member picker group chat.
Tidak mengekspos membership dari sekolah lain.

```json
{
  "members": [
    {
      "userId": "uuid",
      "fullName": "Budi Santoso",
      "email": "budi@siswa.sch.id",
      "roles": ["student"]
    }
  ]
}
```

### Open School Room

`POST /school/open`

Membuka atau membuat room utama untuk active school.

```json
{
  "room": {
    "roomId": "uuid",
    "roomName": "Ruang sekolah",
    "roomType": "group",
    "roomRefType": "school",
    "roomRefId": "school-uuid",
    "schoolId": "school-uuid",
    "schoolName": "SMA EduVerse",
    "unreadCount": 0,
    "canSend": true
  }
}
```

### Create Group Room

`POST /groups`

```json
{
  "roomName": "Grup Belajar Fisika",
  "memberUserIds": ["user-uuid"]
}
```

Rules:

- Current user harus active school member.
- Semua `memberUserIds` harus active school member pada active school.
- Duplicate member ditolak.
- Creator selalu dimasukkan sebagai room member dengan `crm_role = "admin"`.
- Member terpilih dimasukkan dengan `crm_role = "member"`.
- `room_ref_type` dan `room_ref_id` disimpan `NULL`.

Response:

```json
{
  "room": {
    "roomId": "uuid",
    "roomName": "Grup Belajar Fisika",
    "roomType": "group",
    "roomRefType": null,
    "roomRefId": null,
    "schoolId": "school-uuid",
    "schoolName": "SMA EduVerse",
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
      "content": "Selamat pagi.",
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
  "content": "Halo semua."
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
