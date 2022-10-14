-- name: CreateAccount :one
Insert into "Accounts" ("owner", "balance", "currency")
Values ($1, $2, $3)
RETURNING *;
-- name: UpdateAccount :exec
UPDATE "Accounts"
set owner = $2,
    balance = $3,
    currency = $4
WHERE id = $1;
-- name: GetAccount :one
SELECT *
FROM "Accounts"
WHERE id = $1
LIMIT 1;
-- name: GetAccountForUpdate :one
SELECT *
FROM "Accounts"
WHERE id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListAccounts :many
SELECT *
FROM "Accounts"
ORDER BY "id"
LIMIT $1 OFFSET $2;
-- name: DeleteAccount :exec
Delete from "Accounts"
where id = $1;