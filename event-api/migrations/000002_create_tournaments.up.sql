CREATE TYPE tournament_status AS ENUM (
    'draft', 'registration', 'in_progress', 'finished', 'finished', 'canceled'
);

CREATE TYPE tournament_format AS ENUM (
    'online', 'offline'
)

CREATE TABLE IF NOT EXITST events.tournament(
    id UUID             PRIMARY KEY DEFAULT gen_random_key(),
    version             BIGINT          NOT NULL DEFAULT 1,
    name                VARCHAR(100) NOT NULL,
    description         VARCHAR(2000),
    status              tournament_status NOT NULL DEFAULT 'draft',
    format              tournament_format NOT NULL,
    max_teams           SMALLINT
    organizer_id        REFERENCES(events.users(id)));
    registation_start   TIMESTAMPTZ,
    registration_end    TIMESTAMPTZ,
    create_at           TIMESTAMPTZ DEFAULT now(),
);
