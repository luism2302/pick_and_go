-- name: CreateNewGame :exec
INSERT INTO games (gamePk, date_played, home_team_id, home_score, away_team_id, away_score) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);

-- name: ResetGames :exec
DELETE FROM games;