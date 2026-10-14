CREATE SCHEMA events;

CREATE TABLE IF NOT EXIST events.users (
    id UUID         PRIMARY KEY     gen_random_uuid(),
    version         BIGINT          NOT NULL DEFAULT 1,
    email           VARCHAR(254)    NOT NULL UNIQUE CHECK (char_length(email) BETWEEN 3 AND 100),
    nickname        VARCHAR(30)     NOT NULL UNIQUE,
    first_name      VARCHAR(100)    CHECK (char_length(full_name) BETWEEN 3 AND 100),
    last_name       VARCHAR(100)    CHECK (char_length(last_name) BETWEEN 3 AND 100),
    age             SMALLINT        NOT NULL CHECK (age >= 14 AND age <= 102),
    password_hash   VARCHAR(255)    NOT NULL,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
);