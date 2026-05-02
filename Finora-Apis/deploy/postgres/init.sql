-- Finora bootstrap schema (runs on first Postgres container start).

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Demo user: email demo@finora.local / password demo123 (change in production).
INSERT INTO users (email, password_hash)
VALUES (
    'demo@finora.local',
    '$2a$10$nRmifKXWOFmoTNtuJ.hIyeTljlZZrdDNQiMn9iXCAr3pkHWN43RMW'
)
ON CONFLICT (email) DO NOTHING;
