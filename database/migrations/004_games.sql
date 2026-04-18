-- +goose Up
CREATE TABLE games (
    gamePk INT PRIMARY KEY,
    date_played DATE NOT NULL,
    home_team_id INT REFERENCES teams(id) ON DELETE CASCADE NOT NULL,
    home_score INT NOT NULL,
    away_team_id INT REFERENCES teams(id) ON DELETE CASCADE NOT NULL,
    away_score INT NOT NULL
);

-- +goose Down
DROP TABLE games;