-- name: CreateTeam :exec
INSERT INTO teams(id, team_name, abbreviation,division_id) VALUES (
    $1,
    $2,
    $3,
    $4
);

-- name: ResetTeams :exec
DELETE FROM teams;