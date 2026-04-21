-- +goose Up
CREATE TABLE players (
    id INT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT NOT NULL,
    is_active BOOLEAN NOT NULL,
    team_id INT REFERENCES teams(id) ON DELETE CASCADE NOT NULL,
    primary_position TEXT NOT NULL,
    batside TEXT NOT NULL,
    pitchhand TEXT NOT NULL
);

-- +goose Down
DROP TABLE players;