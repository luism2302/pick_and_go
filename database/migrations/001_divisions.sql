-- +goose Up
CREATE TABLE divisions (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE divisions;