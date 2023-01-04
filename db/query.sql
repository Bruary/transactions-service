-- name: GetTransaction :one
SELECT * FROM transactions
WHERE uid = $1 LIMIT 1;

-- name: GetTransactions :many
SELECT * FROM transactions
ORDER BY id;

-- name: InsertTransaction :one
INSERT INTO transactions (
  amount, currency
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE uid = $1;

-- name: UpdateTransaction :exec
UPDATE transactions
  set amount = $2,
  currency = $3
WHERE uid = $1;