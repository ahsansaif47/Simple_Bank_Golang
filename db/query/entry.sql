-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
Returning *;
-- name: GetEntry :one
SELECT *
from entries
where id = $1
LIMIT 1;
-- name: ListEntries :many
SELECT *
from entries
where account_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;