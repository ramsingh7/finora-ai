# Finora Web Admin Panel

Production-ready frontend boilerplate for Finora admin panel built with:

- React + TypeScript + Vite
- Tailwind CSS
- Redux Toolkit + React Redux
- React Router
- Axios + Zod + React Hook Form

## Getting started

```bash
npm install
cp .env.example .env
npm run dev
```

## Scripts

- `npm run dev` - start local development server
- `npm run lint` - run ESLint
- `npm run build` - run type-check and production build
- `npm run preview` - preview production build locally

## Implemented phase-1 modules

- `/login`
- `/dashboard`
- `/dashboard/system-health`
- `/dashboard/profile`
- `/dashboard/settings`
- Placeholder routes for users, roles, market-data, models, predictions, jobs and logs

## API expectations

- `POST /api/auth/login` -> `{ ok: true, expiresInSeconds }`
- `GET /api/health` -> `{ status, version }`

`POST /api/auth/logout` is expected for backend logout endpoint and can be wired in `handleLogout`.
