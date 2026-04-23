-- name: CreateDivision :exec
INSERT INTO Divisions (id, name) VALUES (
    $1,
    $2
);

-- name: ResetDivisions :exec
DELETE FROM divisions;
