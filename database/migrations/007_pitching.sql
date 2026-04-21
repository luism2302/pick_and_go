-- +goose Up
CREATE TABLE pitching (
    id UUID PRIMARY KEY,
    player_id INT REFERENCES players(id) ON DELETE CASCADE NOT NULL,
    gamePk INT REFERENCES games(gamePk) ON DELETE CASCADE NOT NULL,
    era TEXT NOT NULL,
    whip TEXT NOT NULL,
    innings_pitched TEXT NOT NULL,
    strikeouts INT NOT NULL,
    walks INT NOT NULL,
    home_runs INT NOT NULL,
    earned_runs INT NOT NULL,
    hits INT NOT NULL,
    wins INT NOT NULL,
    losses INT NOT NULL, 
    games_started INT NOT NULL,
    saves INT NOT NULL,
    blown_saves INT NOT NULL,
    strikeouts_per9 TEXT NOT NULL,
    walks_per9 TEXT NOT NULL,
    strikeout_walk_ratio TEXT NOT NULL
);

-- +goose Down
DROP TABLE pitching;