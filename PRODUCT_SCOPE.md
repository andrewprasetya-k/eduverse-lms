# Eduverse LMS Product Scope

## 1. Product Vision

Eduverse LMS is a multi-school learning management system for managing academic classes, learning content, assignments, submissions, grading, and school communication in one consistent platform.

The product should help schools run day-to-day digital learning workflows with clear role boundaries, reliable academic records, and a simple experience for teachers, students, and school administrators.

## 2. Target Users

- School administrators who manage academic structure, users, classes, and school-level configuration.
- Teachers who publish learning content, create assignments, assess submissions, and communicate with class members.
- Students who access materials, submit assignments, track progress, and receive notifications.
- System administrators who manage multi-school access, platform setup, and operational health.

## 3. User Roles

- Super Admin: Manages platform-level access and school setup. Does not automatically perform school academic actions unless assigned a school role.
- Admin: Manages school-level academic structure, users, classes, enrollments, and operational data.
- Teacher: Manages assigned classes, materials, assignments, assessments, comments, and class communication.
- Student: Accesses class content, submits work, views grades, comments, and receives notifications.

## 4. Core Workflows

- School setup: create school structure, academic years, terms, subjects, classes, and class enrollments.
- Class learning: teachers publish materials and students consume content by class context.
- Assignment lifecycle: teachers create assignments, students submit work, teachers assess submissions, and students view outcomes.
- Grade management: configure assessment weights, calculate grades, and view student/class grade reports.
- Communication: feeds and comments support class-level discussion and contextual interaction.
- Notifications: users receive activity updates for important academic events.
- Media management: users upload and attach files to learning objects through controlled storage.

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

## 7. Features Intentionally Out of Scope

- Full HR, payroll, finance, billing, or accounting modules.
- Parent portal and guardian-facing workflows.
- Video conferencing or live classroom streaming.
- Complex exam proctoring or anti-cheating systems.
- Marketplace, ecommerce, or paid course management.
- Offline-first mobile synchronization.
- Social networking features beyond class learning communication.
- Deep SIS/ERP integrations unless explicitly prioritized later.

## 8. Technical Principles

- Keep academic data ownership explicit by school, class, subject class, and user role.
- Prefer clear backend authorization over frontend-only access control.
- Preserve auditability and avoid destructive data loss for academic records.
- Keep API contracts stable and documented before frontend wiring.
- Build incrementally with small, verifiable changes.
- Treat file storage and media lifecycle as first-class product behavior, not placeholder metadata.
- Fail safely for authentication, authorization, storage, and data integrity errors.

## 9. Frontend Priorities

- Build role-specific experiences for admin, teacher, and student workflows.
- Prioritize core learning flows before advanced analytics or customization.
- Make school and class context visible and hard to confuse.
- Keep assignment submission, grading, material browsing, and notifications easy to test end to end.
- Use backend permissions as the source of truth while using frontend gating for usability.
- Present file upload, validation errors, and storage failures clearly.
- Keep UI patterns consistent across CRUD, lists, detail pages, and forms.

## 10. Backend Priorities

- Maintain reliable RBAC and school membership enforcement.
- Keep route ordering and route parameters unambiguous.
- Keep actor identity derived from JWT, not trusted request body fields.
- Protect academic data integrity when deleting schools, classes, subjects, materials, assignments, and media.
- Keep notification triggers best-effort unless a workflow explicitly requires hard failure.
- Continue documenting API behavior as contracts evolve.
- Add focused regression tests around auth, RBAC, route conflicts, storage lifecycle, and academic delete protections.

## 11. Future Scalability Considerations

- Support larger schools with pagination, search, filtering, and efficient list endpoints.
- Keep storage provider abstractions portable for future providers beyond Supabase.
- Design notification delivery so in-app, email, and realtime channels can evolve independently.
- Prepare reporting/export flows for background jobs when data volume grows.
- Keep multi-school boundaries strict for data isolation and future tenant-level configuration.
- Consider caching only after correctness and authorization rules are stable.
- Keep schema changes migration-friendly as academic workflows expand.
