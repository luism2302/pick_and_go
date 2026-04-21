-- name: CreateNewPlayer :exec
INSERT INTO players(id, first_name, last_name, age, is_active, team_id, primary_position, batside, pitchhand) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
);

-- name: ResetPlayers :exec
DELETE FROM players;