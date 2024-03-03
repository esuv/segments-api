CREATE TABLE IF NOT EXISTS USERS
(
    id             SERIAL PRIMARY KEY,
    username       VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS SEGMENTS
(
    id     SERIAL PRIMARY KEY,
    name   VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS USER_SEGMENTS
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER,
    segment_id     INTEGER,
    created        TIMESTAMP NOT NULL
    expired        TIMESTAMP
);

CREATE TABLE IF NOT EXISTS HISTORY
(
    id             SERIAL PRIMARY KEY,
    user_id        INTEGER NOT NULL,
    segment_id     INTEGER NOT NULL,
    operation      VARCHAR(255) NOT NULL,
    date           TIMESTAMP NOT NULL
);

