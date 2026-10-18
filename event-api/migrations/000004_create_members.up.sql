CREATE TYPE events.team_member_role AS ENUM ('captain', 'player');

CREATE TABLE IF NOT EXISTS events.team_members (
    team_id       UUID NOT NULL,
    tournament_id UUID NOT NULL,
    user_id       UUID NOT NULL REFERENCES events.users(id) ON DELETE CASCADE,
    role          events.team_member_role NOT NULL DEFAULT 'player',
    joined_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (team_id, user_id),
    FOREIGN KEY (team_id, tournament_id) REFERENCES events.teams(id, tournament_id) ON DELETE CASCADE,
    UNIQUE (tournament_id, user_id) -- Гарантирует: 1 юзер = максимум 1 команда в рамках одного турнира
);