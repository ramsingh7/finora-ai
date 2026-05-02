# Finora Admin Panel Frontend Guide

This document is a practical blueprint to build an admin panel frontend for the current Finora backend in this repository.

## 1) Current Backend Status (Important)

Your backend currently exposes these gRPC methods in `finora.v1.FinoraService`:

- `Health(HealthRequest) -> HealthResponse`
- `Login(LoginRequest) -> LoginResponse`

That means:

- Authentication can be implemented now.
- Most admin screens (users, analytics, models, jobs, alerts, etc.) will need new backend APIs first.

Use this guide to build the admin panel foundation now, then plug in new APIs as they are added.

## 2) Recommended Frontend Stack

- Framework: `Next.js` (App Router, TypeScript)
- UI: `Tailwind CSS` + `shadcn/ui`
- Data fetching: `TanStack Query`
- Forms: `react-hook-form` + `zod`
- State (lightweight): `zustand` (optional)
- Charts: `recharts`
- Auth/session: HTTP-only cookies (recommended via BFF)

## 3) Architecture Recommendation

## Prefer BFF (Backend-for-Frontend) for Admin

Because browsers do not natively work well with raw gRPC, create a BFF layer in Next.js:

- Browser -> Next.js Route Handler (`/api/*`) via HTTP/JSON
- Next.js server -> gRPC backend (`finora.v1.FinoraService`)

Benefits:

- No token leakage in browser storage.
- Easier RBAC checks and request validation.
- Cleaner migration when more gRPC methods are added.

## Runtime Flow

1. Admin enters email/password on login page.
2. Next.js route calls backend `Login`.
3. Next.js stores JWT in secure HTTP-only cookie.
4. Frontend pages call internal `/api/*` endpoints.
5. BFF reads cookie, injects auth metadata to gRPC calls.

## 4) Minimum Admin Panel Scope (Phase 1)

Implement these first:

- Login page
- Dashboard shell (sidebar, topbar, profile menu)
- Health status widget (from backend `Health`)
- Session management (login/logout/protected routes)
- Error boundaries and global notifications
- Audit-ready layout for future modules

## 5) Suggested Routes

- `/login`
- `/dashboard`
- `/dashboard/system-health`
- `/dashboard/profile`
- `/dashboard/settings`

Future-ready placeholders:

- `/dashboard/users`
- `/dashboard/roles`
- `/dashboard/market-data`
- `/dashboard/models`
- `/dashboard/predictions`
- `/dashboard/jobs`
- `/dashboard/logs`

## 6) Suggested Project Structure (Frontend Repo)

```text
finora-admin/
  src/
    app/
      (auth)/
        login/page.tsx
      (protected)/
        dashboard/layout.tsx
        dashboard/page.tsx
        dashboard/system-health/page.tsx
      api/
        auth/login/route.ts
        auth/logout/route.ts
        health/route.ts
    components/
      layout/
      ui/
      charts/
    features/
      auth/
      dashboard/
      health/
    lib/
      grpc/
      api-client/
      env.ts
      auth.ts
      validators.ts
    hooks/
    types/
```

## 7) API Contract Mapping (Current)

## gRPC to UI Mapping

- `Health`
  - UI usage: system status card
  - Return fields used: `status`, `version`
- `Login`
  - UI usage: login form submit
  - Return fields used: `access_token`, `expires_in_seconds`

## BFF Endpoints to Create

- `POST /api/auth/login`
  - Input: `{ email, password }`
  - Action: call gRPC `Login`, set secure cookie
  - Output: `{ ok: true, expiresInSeconds }`

- `POST /api/auth/logout`
  - Action: clear cookie
  - Output: `{ ok: true }`

- `GET /api/health`
  - Action: call gRPC `Health`
  - Output: `{ status, version }`

## 8) Auth and Security Checklist

- Use HTTP-only, secure, same-site cookies for access token.
- Do not store JWT in `localStorage`.
- Add route-level auth guard for all `/dashboard/*`.
- Add server-side validation using `zod`.
- Add request timeout and retry policy for BFF -> gRPC calls.
- Log failed logins and suspicious request bursts.

## 9) UI/UX Guidelines for Admin Panel

- Keep sidebar persistent on desktop, collapsible on tablet/mobile.
- Use card-based dashboard with high signal metrics.
- Add skeleton loaders and empty states.
- Show clear error states (network/auth/server).
- Include dark mode and accessibility (keyboard + contrast).

## 10) Environment Variables (Frontend)

For Next.js admin app:

```env
NODE_ENV=development
FINORA_GRPC_ADDR=localhost:50051
FINORA_API_TIMEOUT_MS=10000
SESSION_COOKIE_NAME=finora_admin_session
SESSION_COOKIE_SECURE=false
```

For production, set:

- `SESSION_COOKIE_SECURE=true`
- `NODE_ENV=production`
- gRPC address to internal service DNS

## 11) Development Plan (Execution Order)

1. Create Next.js + TypeScript project.
2. Add Tailwind + shadcn + base layout.
3. Build login page + form validation.
4. Create `/api/auth/login` and `/api/auth/logout`.
5. Add cookie session helpers and protected route middleware.
6. Build dashboard shell and health card.
7. Create `/api/health` and connect to UI.
8. Add global toast, error boundary, loading skeleton.
9. Add unit tests for auth and health flows.
10. Add CI checks (lint, typecheck, test, build).

## 12) Testing Checklist

- Login success with valid credentials.
- Login rejection with wrong credentials.
- Protected route redirect when not authenticated.
- Health card loads status + version.
- Cookie expiry forces re-login.
- Network failure shows recoverable UI error.

## 13) Future Backend APIs Needed for Full Admin Panel

To complete a real admin panel, backend should add methods for:

- User management (list, create, disable, role update)
- Role and permission management (RBAC)
- System metrics and logs
- Model registry and model deployment status
- Prediction job queues and job history
- Data source sync status and scheduler controls

---

If you want, I can generate the initial `finora-admin` Next.js starter code next (with login page, protected dashboard layout, and health widget wired to your current backend).
