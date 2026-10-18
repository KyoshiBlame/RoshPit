CREATE TABLE IF NOT EXISTS events.teams (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version       BIGINT NOT NULL DEFAULT 1,
    tournament_id UUID NOT NULL REFERENCES events.tournaments(id) ON DELETE CASCADE,
    name          VARCHAR(50) NOT NULL CHECK (char_length(name) BETWEEN 3 AND 50),
    captain_id    UUID REFERENCES events.users(id) ON DELETE SET NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tournament_id, name),
    UNIQUE (id, tournament_id) -- Нужно для составного FK в team_members
);