CREATE TYPE match_status AS ENUM (
    'scheduled', 'in_progress', 'paused', 'finished', 'cancelled'
);

CREATE TABLE IF NOT EXISTS events.matches (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    version       BIGINT       NOT NULL DEFAULT 1,
    tournament_id UUID         NOT NULL REFERENCES events.tournaments(id) ON DELETE CASCADE,
    team_a_id     UUID         NOT NULL REFERENCES events.teams(id),
    team_b_id     UUID         NOT NULL REFERENCES events.teams(id),
    judge_id      UUID         REFERENCES events.users(id) ON DELETE SET NULL,
    status        match_status NOT NULL DEFAULT 'scheduled',
    round         SMALLINT     NOT NULL DEFAULT 1 CHECK (round > 0),
    scheduled_at  TIMESTAMPTZ,
    started_at    TIMESTAMPTZ,
    finished_at   TIMESTAMPTZ,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at    TiMESTAMPTZ  NOT NULL DEFAULT now(),
    CHECK (team_a_id != team_b_id)
);

CREATE TABLE IF NOT EXITS