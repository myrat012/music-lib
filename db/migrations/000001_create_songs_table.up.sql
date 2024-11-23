CREATE TABLE songs (
    id           SERIAL PRIMARY KEY,
    group_song   VARCHAR(255) NOT NULL,
    song         VARCHAR(255) NOT NULL,
    release_date DATE             NULL,
    text         TEXT             NULL,
    link         VARCHAR(255)     NULL
);