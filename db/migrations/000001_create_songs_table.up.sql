CREATE TYPE state_t AS ENUM ('enabled', 'disabled', 'deleted');

CREATE TABLE songs (
    id           SERIAL PRIMARY KEY,
    group        VARCHAR(255) NOT NULL,
    song         VARCHAR(255) NOT NULL,
    release_date DATE         NOT NULL,
    text         TEXT             NULL,
    link         VARCHAR(255) NOT NULL,
    state        state_t      NOT NULL,
);