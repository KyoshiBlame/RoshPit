CRATE TABLE IF NOT EXIST events.teams (
    id              UUID            PRIMARY KEY DEFAULT gen_random(),
    version         BIGINT          NOT NULL DEFAULT 1,
    tounament_id    UUID            NOT NULL REFERENCES tournament(id) ON DELETE CASCADE,
    name            VARCHAR(50)     NOT NULL UNIQUE CHEK(char_length(name) BETWEEN 3 AND 50),
    captain_id      UUID            REFERENCES user(id) ON DELETE SET NULL,
    created_at      TIMESTAMPTZ     NOT NULL DEFAUL now(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAUL now(),
)