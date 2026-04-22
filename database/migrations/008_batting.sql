-- +goose Up
CREATE TABLE batting(
    id UUID PRIMARY KEY,
    player_id INT REFERENCES players(id) ON DELETE CASCADE NOT NULL,
    gamePk INT REFERENCES games(gamePk) ON DELETE CASCADE NOT NULL,
    at_bats INT NOT NULL,
    runs INT NOT NULL,
    hits INT NOT NULL,
    doubles INT NOT NULL,
    triples INT NOT NULL,
    home_runs INT NOT NULL,
    rbi INT NOT NULL,
    stolen_bases INT NOT NULL,
    caught_stealing INT NOT NULL,
    walks INT NOT NULL,
    strikeouts INT NOT NULL,
    hit_by_pitch INT NOT NULL,
    avg TEXT NOT NULL,
    obp TEXT NOT NULL,
    slugging TEXT NOT NULL,
    ops TEXT NOT NULL,
    left_on_base INT NOT NULL,
    sac_bunts INT NOT NULL,
    sac_flies INT NOT NULL
);

-- +goose Down
DROP TABLE batting;