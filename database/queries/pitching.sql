-- name: CreatePitchingEntry :exec
INSERT INTO pitching (id, player_id, gamePk, era, whip, innings_pitched, strikeouts, walks, home_runs, earned_runs, hits, wins, losses, games_started, saves, blown_saves, strikeouts_per9, walks_per9, strikeout_walk_ratio) VALUES (
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
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18
);

-- name: ResetPitching :exec
DELETE FROM pitching;