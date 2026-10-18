CREATE TYPE events.tournament_status AS ENUM (
    'draft', 'registration', 'in_progress', 'finished', 'cancelled'
);

CREATE TYPE events.tournament_format AS ENUM (
    'online', 'offline'
);

CREATE TABLE IF NOT EXISTS events.tournaments (
    id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version            BIGINT NOT NULL DEFAULT 1,
    name               VARCHAR(100) NOT NULL,
    description        VARCHAR(2000),
    status             events.tournament_status NOT NULL DEFAULT 'draft',
    format             events.tournament_format NOT NULL,
    max_teams          SMALLINT NOT NULL CHECK (max_teams >= 2),
    organizer_id       UUID NOT NULL REFERENCES events.users(id),
    registration_start TIMESTAMPTZ,
    registration_end   TIMESTAMPTZ,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);