-- name: CreateTransfer :one
Insert into "Transfers" (
        "from_account",
        "to_account",
        "transaction_time",
        "amount"
    )
Values ($1, $2, $3, $4)
RETURNING *;
-- name: GetTransfer :one
SELECT *
from "Transfers"
WHERE id = $1
LIMIT 1;
-- name: ListTransfers :many
SELECT *
from "Transfers"
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateTransfer :exec
UPDATE "Transfers"
SET from_account = $2,
    to_account = $3,
    transaction_time = $4,
    amount = $5
WHERE id = $1;
-- name: DeleteTransfer :one
Delete from "Transfers"
WHERE id = $1
RETURNING *;