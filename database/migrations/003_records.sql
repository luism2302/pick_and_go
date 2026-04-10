-- +goose Up
CREATE TABLE records (
   id UUID PRIMARY KEY,
   team_id INT REFERENCES teams(id) ON DELETE CASCADE NOT NULL,
   wins INT NOT NULL,
   losses INT NOT NULL,
   pct TEXT NOT NULL,
   streak TEXT NOT NULL,
   runs_scored INT NOT NULL,
   runs_against INT NOT NULL,
   home_wins INT NOT NULL, 
   home_losses INT NOT NULL,
   away_wins INT NOT NULL,
   away_losses INT NOT NULL
);

-- +goose Down
DROP TABLE records;