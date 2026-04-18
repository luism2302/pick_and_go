-- name: CreateNewInning :exec
INSERT INTO innings(id, gamePk, inning_name, home_runs, home_hits, home_errors, away_runs, away_hits, away_errors) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);

-- name: ResetInnings :exec
DELETE FROM innings;