# 🧠 Project Handoff Context: Eduverse LMS Backend

## 📌 Project Overview
Eduverse is a Learning Management System (LMS) built with **Go (Gin Framework)** and **GORM**. The system uses a multi-school architecture where users can belong to one or more schools.

## 🏗️ Architectural Patterns (Strictly Follow These)
1.  **Layered Architecture**:
    *   `domain`: DB models and TableName definitions.
    *   `dto`: Request/Response structures (JSON).
    *   `repository`: Raw GORM queries.
    *   `service`: Business logic, validation, and cross-service coordination.
    *   `handler`: HTTP request parsing and DTO mapping.

2.  **ID vs Code Convention**:
    *   **Internal (DB)**: Use UUID for all Primary and Foreign Keys.
    *   **External (URL)**: Use `schoolCode` or `subjectCode` in API paths for human-readability.
    *   **Converter**: Always use `schoolService.ConvertCodeToID(code)` in the Service layer to translate codes to internal IDs.

3.  **Data Integrity**:
    *   Repositories **MUST** check `RowsAffected == 0` on Update/Delete/Patch and return `gorm.ErrRecordNotFound` if no row was modified.
    *   Services **MUST** handle these errors and return user-friendly messages.

4.  **API Standards**:
    *   Standardized `SchoolHeaderDTO` (ID, Name, Code, Logo) used across all modules when returning school context.
    *   Use `Preload` in Repositories to ensure related data (like Creator names or School info) is included in responses.
    *   All activation/deactivation actions use the `PATCH` method.

## 🔐 Security
*   Passwords are hashed using **Bcrypt**.
*   RBAC system is implemented: `User` -> `SchoolUser` -> `UserRole` -> `Role` -> `Permission`.
*   *Note: Auth middleware (JWT) is NOT yet implemented.*

## 📂 Documentation
Full API specs are available in `backend/docs/api/`. Refer to these files before changing any endpoint behavior.

## 🛠️ Tech Stack Info
*   **Database**: PostgreSQL (Supabase).
*   **Env**: Managed via `.env` (loaded in `main.go`).
*   **Build**: Run `go build ./...` to verify all modules.

---
*Generated on: 13-02-2026. Use this as a prompt for future turns.*
