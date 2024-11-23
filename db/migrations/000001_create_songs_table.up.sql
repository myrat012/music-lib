CREATE TABLE songs (
    id           SERIAL PRIMARY KEY,
    group_song   VARCHAR(255) NOT NULL,
    song         VARCHAR(255) NOT NULL,
    release_date VARCHAR(255)     NULL,
    text         TEXT             NULL,
    link         VARCHAR(255)     NULL
);