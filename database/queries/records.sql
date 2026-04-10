-- name: CreateTeamRecord :exec
INSERT INTO records(id, team_id, wins, losses, pct, streak, runs_scored, runs_against, home_wins, home_losses, away_wins, away_losses) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
);