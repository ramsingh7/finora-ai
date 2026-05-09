-- Migration 001: rename users → mst_users
-- Safe to re-run: checks existence before acting.

DO $$
BEGIN
    -- Rename table if it still has the old name
    IF EXISTS (SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = 'users')
       AND NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = 'mst_users')
    THEN
        ALTER TABLE users RENAME TO mst_users;
    END IF;
END;
$$;
