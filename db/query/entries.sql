-- name: CreateEntry :one
Insert into "Entries" (
        "account_id",
        "amount",
        "created_at"
    )
Values ($1, $2, $3)
RETURNING *;
-- name: GetEntry :one
SELECT *
from "Entries"
WHERE id = $1
LIMIT 1;
-- name: ListEntries :many
SELECT *
from "Entries"
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateEntry :exec
UPDATE "Entries"
set account_id = $2,
    amount = $3,
    created_at = $4
WHERE id = $1;
-- name: DeleteEntry :one
Delete from "Entries"
WHERE id = $1
RETURNING *;