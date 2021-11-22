-- name: GetUser :one
SELECT
	*
FROM
	users
WHERE
	id = ?;

-- name: ListUsers :many
SELECT
	*
FROM
	users
ORDER BY
	name;

-- name: CreateUser :execresult
INSERT INTO
	users(name, email)
VALUES
	(?, ?);

-- name: UpdateUser :exec
UPDATE
	users
SET
	name = ?,
	email = ?
WHERE
	id = ?;

-- name: DeleteUser :exec
DELETE FROM
	users
WHERE
	id = ?;