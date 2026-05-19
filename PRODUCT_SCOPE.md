# Eduverse LMS Product Scope

## 1. Product Vision

Eduverse LMS is a multi-school learning management system for managing academic classes, learning content, assignments, submissions, grading, school communication, real-time chat, and student notes in one consistent platform.

The product should help schools run day-to-day digital learning workflows with clear role boundaries, reliable academic records, real-time class communication, and a simple experience for teachers, students, and school administrators.

## 2. Target Users

- School administrators who manage academic structure, users, classes, and school-level configuration.
- Teachers who publish learning content, create assignments, assess submissions, and communicate with class members.
- Students who access materials, submit assignments, track progress, take notes, and receive notifications.
- System administrators who manage multi-school access, platform setup, and operational health.

## 3. User Roles

- Super Admin: Manages platform-level access and school setup. Does not automatically perform school academic actions unless assigned a school role.
- Admin: Manages school-level academic structure, users, classes, enrollments, and operational data.
- Teacher: Manages assigned classes, materials, assignments, assessments, comments, class communication, and chat channels.
- Student: Accesses class content, submits work, views grades, comments, takes notes per material, participates in class chat, and receives notifications.

## 4. Core Workflows

- School setup: create school structure, academic years, terms, subjects, classes, and class enrollments.
- Class learning: teachers publish materials and students consume content by class context.
- Assignment lifecycle: teachers create assignments, students submit work, teachers assess submissions, and students view outcomes.
- Grade management: configure assessment weights, calculate grades, and view student/class grade reports.
- Communication: feeds and comments support class-level discussion and contextual interaction.
- Real-time chat: class and subject channels for live communication between teachers and students, with DM support for private teacher-student conversation.
- Student notes: personal note-taking per material, tied to academic context, accessible anytime during learning.
- Notifications: users receive activity updates for important academic events.
- Media management: users upload and attach files to learning objects and chat messages through controlled storage.

## 5. MVP Features

- Multi-school user and role model.
- JWT authentication and role-based access control.
- School, term, subject, class, subject-class, and enrollment management.
- Material creation, listing, detail, update, delete, and attachment support.
- Assignment creation, submission, assessment, status tracking, and attachment support.
- Grade book with weighted assessment categories and grade reports.
- Notifications for major learning events.
- Feed and comment features for learning communication.
- Media upload and storage integration for files.
- Basic dashboards for student, teacher, and admin contexts.

## 6. Post-MVP Features

### Academic
- Assignment extension request and approval flow.
- Signed/private download URLs for protected media access.
- Thumbnail generation and richer media previews.
- Transcript and report card export.
- Rich text content authoring with sanitization.
- Nested comments and improved discussion threads.
- Unified activity timeline across materials, assignments, feeds, comments, and grades.
- Class schedule and timetable management.
- Material progress analytics and engagement reporting.
- Notification preferences and optional email delivery.

### Real-time Chat
- Class channel: one real-time chat room per class, auto-created on class creation, auto-populated from enrollments.
- Subject channel: one real-time chat room per subject class, auto-created on subject class creation.
- Direct message: private one-on-one chat between any two users (student ↔ student, teacher ↔ student, teacher ↔ teacher).
- Free group chat: any user can create a custom group chat with any members within the same school, independent of academic structure.
- Thread/reply per message.
- File sharing in chat via existing media system.
- Mention (@user) support.
- Read receipts (per-message read status).
- Typing indicators.
- Online/offline presence.
- Message unsend (soft delete).
- Academic context linking: teachers can share links to materials or assignments directly from chat with rendered preview cards.
- System messages for academic events (new material added, new assignment created).
- Pin important messages (teacher only for academic channels; creator/admin for free groups).

### Student Notes
- Personal note per material: each student has one private note per material, scoped to their enrollment context.
- Rich text / markdown support.
- Auto-save on edit.
- Notes accessible from material detail page.

## 7. Technical Architecture

### Backend
- Language: Go
- Framework: Gin
- ORM: GORM
- Database: PostgreSQL (Supabase)
- Auth: JWT
- Real-time: Gorilla WebSocket
- File storage: Supabase Storage / S3-compatible

### Frontend
- Framework: Vue 3 + Vite
- Routing: Vue Router
- State Management: Pinia
- HTTP Client: Axios
- Real-time: Native WebSocket (Gorilla-compatible)
- Styling: Tailwind CSS

### Folder Structure (Backend)
```
internal/
├── domain/        ← GORM models
├── dto/           ← request/response contracts
├── repository/    ← database queries
├── service/       ← business logic
├── handler/
│   ├── rest/      ← Gin HTTP handlers
│   └── ws/        ← WebSocket hub & handlers (chat)
└── middleware/    ← auth, RBAC
```

### Folder Structure (Frontend)
```
src/
├── assets/        ← static assets
├── components/    ← reusable Vue components
├── composables/   ← reusable logic (useAuth, useChat, dll)
├── layouts/       ← layout per role (admin, teacher, student)
├── pages/         ← route-based views
├── router/        ← Vue Router config
├── stores/        ← Pinia stores
└── services/      ← Axios API calls
```

### Chat Architecture
- WebSocket Hub manages all active connections in memory.
- Each chat room runs as a goroutine with its own broadcast channel.
- Participants derived from enrollments and synced to `chat_room_members` on enrollment events.
- Read receipts, typing indicators, and online/offline presence handled via WebSocket events.
- Chat rooms auto-created when a class or subject class is created.
- Chat room members auto-populated when a student is enrolled.

### Database Schema Additions (Chat & Notes)
- `chat_rooms`: linked to `classes` or `subject_classes` via `room_ref_type` + `room_ref_id`.
- `chat_room_members`: explicit member table, populated from `enrollments`, supports DM.
- `chat_messages`: supports text, file, system message types, reply threads, and academic context references.
- `chat_attachments`: references existing `medias` table.
- `chat_read_receipts`: tracks last read message per user per room.
- `student_notes`: one note per student per material, rich text content.
- `classes` table: add `cls_chat_room_id` pointer to chat room.
- `subject_classes` table: add `scl_chat_room_id` pointer to chat room.

## 8. Features Intentionally Out of Scope

- Full HR, payroll, finance, billing, or accounting modules.
- Parent portal and guardian-facing workflows.
- Video conferencing or live classroom streaming.
- Complex exam proctoring or anti-cheating systems.
- Marketplace, ecommerce, or paid course management.
- Offline-first mobile synchronization.
- Social networking features beyond class learning communication.
- Deep SIS/ERP integrations unless explicitly prioritized later.

## 9. Technical Principles

- Keep academic data ownership explicit by school, class, subject class, and user role.
- Prefer clear backend authorization over frontend-only access control.
- Preserve auditability and avoid destructive data loss for academic records.
- Keep API contracts stable and documented before frontend wiring.
- Build incrementally with small, verifiable changes.
- Treat file storage and media lifecycle as first-class product behavior, not placeholder metadata.
- Fail safely for authentication, authorization, storage, and data integrity errors.
- Keep WebSocket and REST handlers separated in codebase for clarity and maintainability.
- Chat room lifecycle (creation, member sync) must follow academic entity lifecycle (class creation, enrollment).

## 10. Frontend Priorities (Vue 3 + Vite)

- Build role-specific experiences for admin, teacher, and student workflows using Vue Router guards.
- Prioritize core learning flows before advanced analytics or customization.
- Make school and class context visible and hard to confuse.
- Keep assignment submission, grading, material browsing, and notifications easy to test end to end.
- Use backend permissions as the source of truth; use Pinia stores for frontend role-based gating.
- Present file upload, validation errors, and storage failures clearly.
- Keep UI patterns consistent across CRUD, lists, detail pages, and forms.
- Chat UI: class/subject channels accessible from class detail page; DM accessible from user profile or class member list.
- Notes UI: note editor embedded in material detail page, auto-save behavior, markdown preview.
- Use composables (`useAuth`, `useChat`, `useWebSocket`) to encapsulate reusable logic cleanly.

## 11. Backend Priorities

- Maintain reliable RBAC and school membership enforcement.
- Keep route ordering and route parameters unambiguous.
- Keep actor identity derived from JWT, not trusted request body fields.
- Protect academic data integrity when deleting schools, classes, subjects, materials, assignments, and media.
- Keep notification triggers best-effort unless a workflow explicitly requires hard failure.
- Continue documenting API behavior as contracts evolve.
- Add focused regression tests around auth, RBAC, route conflicts, storage lifecycle, and academic delete protections.
- WebSocket Hub must handle concurrent connections safely using goroutines and channels.
- Chat room and member sync must be triggered by enrollment and class creation events, not manual API calls from frontend.

## 12. Development Roadmap

### Phase 1 — Complete MVP Backend (Current)
- [ ] File upload integration (Supabase Storage / S3)
- [ ] Notification triggers integration across all services
- [ ] Assignment extension request and approval flow

### Phase 2 — Post-MVP Academic Features
- [ ] Nested comments
- [ ] Rich text support with sanitization
- [ ] Activity feed / timeline
- [ ] Grade report and transcript export

### Phase 3 — Real-time Chat
- [ ] WebSocket Hub implementation (Gorilla)
- [ ] Class and subject channel creation on academic entity creation
- [ ] Member sync from enrollments
- [ ] Message send, receive, reply, unsend
- [ ] Read receipts and typing indicators
- [ ] File sharing in chat
- [ ] DM (teacher ↔ student)
- [ ] Academic context linking (share material/assignment in chat)

### Phase 4 — Student Notes
- [ ] Note CRUD per material per student
- [ ] Rich text / markdown support
- [ ] Auto-save

### Phase 5 — Frontend (Vue 3 + Vite)
- [ ] Vue + Vite setup with Vue Router and Pinia
- [ ] Role-based routing and layout per role (admin, teacher, student)
- [ ] Admin, teacher, student dashboards
- [ ] All core learning flows (materials, assignments, submissions, grades)
- [ ] Chat UI integrated with native WebSocket
- [ ] Notes UI embedded in material detail

## 13. Future Scalability Considerations

- Support larger schools with pagination, search, filtering, and efficient list endpoints.
- Keep storage provider abstractions portable for future providers beyond Supabase.
- Design notification delivery so in-app, email, and realtime channels can evolve independently.
- Prepare reporting/export flows for background jobs when data volume grows.
- Keep multi-school boundaries strict for data isolation and future tenant-level configuration.
- Consider caching only after correctness and authorization rules are stable.
- Keep schema changes migration-friendly as academic workflows expand.
- WebSocket Hub should be designed for horizontal scaling with Redis pub/sub if traffic grows.