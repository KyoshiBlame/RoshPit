CREATE TYPE team_member_role AS ENUM ('captain', 'player');

CREATE TABLE IF NOT EXIST events.team_members (
    team id     UUID                NOT NULL REFERENCES events.teams(id) ON DELETE CASCADE,
    user_id     UUID                NOT NULL REFERENCES events.users(id) ON DELETE CASCADE,
    role        team_member_role    NOT NULL DEFAULT 'player',
    joined_at   TIMESTAMPTZ         NOT NULL DEFAULT now(),
    PRIMARY KEY (team_id, user_id)
)

CREATE UNIQUE INDEX idx_one_team_per_tournament ON events.team_members (
    user_id,
    (SELECT events.tournament_id FROM events.teams WHERE id = team_id)
);