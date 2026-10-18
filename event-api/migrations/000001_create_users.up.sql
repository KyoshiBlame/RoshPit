CREATE SCHEMA IF NOT EXISTS events;

CREATE TYPE events.user_role AS ENUM ('viewer', 'judge', 'admin');

CREATE TABLE IF NOT EXISTS events.users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version       BIGINT NOT NULL DEFAULT 1,
    email         VARCHAR(254) NOT NULL UNIQUE CHECK (char_length(email) BETWEEN 3 AND 254),
    nickname      VARCHAR(30) NOT NULL UNIQUE CHECK (char_length(nickname) BETWEEN 3 AND 30),
    first_name    VARCHAR(100) CHECK (char_length(first_name) BETWEEN 2 AND 100),
    last_name     VARCHAR(100) CHECK (char_length(last_name) BETWEEN 2 AND 100),
    age           SMALLINT NOT NULL CHECK (age >= 14 AND age <= 102),
    role          events.user_role NOT NULL DEFAULT 'viewer',
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);