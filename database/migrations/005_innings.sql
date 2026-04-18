-- +goose Up
CREATE TABLE innings(
    id UUID PRIMARY KEY,
    gamePk INT REFERENCES games(gamePk) ON DELETE CASCADE NOT NULL,
    inning_name TEXT NOT NULL,
    home_runs INT NOT NULL,
    home_hits INT NOT NULL,
    home_errors INT NOT NULL,
    away_runs INT NOT NULL,
    away_hits INT NOT NULL,
    away_errors INT NOT NULL
);

-- +goose Down
DROP TABLE innings;