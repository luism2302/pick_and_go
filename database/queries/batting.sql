-- name: CreateBattingEntry :exec
INSERT INTO batting(id, player_id, gamePk, at_bats, runs, hits, doubles, triples, home_runs, rbi, stolen_bases, caught_stealing, walks, strikeouts, hit_by_pitch, avg, obp, slugging, ops, left_on_base, sac_bunts, sac_flies) VALUES (
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
    $18,
    $19,
    $20,
    $21
);

-- name: ResetBatting :exec
DELETE FROM batting;