CREATE TABLE IF NOT EXISTS events.tournament_judges (
    tournament_id UUID NOT NULL REFERENCES events.tournaments(id) ON DELETE CASCADE,
    user_id       UUID NOT NULL REFERENCES events.users(id) ON DELETE CASCADE,
    assigned_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (tournament_id, user_id)
);