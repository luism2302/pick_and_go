-- +goose Up
CREATE TABLE teams (
    id INT PRIMARY KEY,
    team_name TEXT UNIQUE NOT NULL,
    abbreviation TEXT UNIQUE NOT NULL,
    division_id INT REFERENCES divisions(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE teams;