-- name: GetAllNodes :many
SELECT * FROM node;

-- name: GetParentNodes :many
SELECT * FROM node WHERE parent_id IS NULL AND owner_id = $1;

-- name: CreateNode :one
INSERT INTO node (id, type, name, parent_id, owner_id)
VALUES ($1, $2, $3, $4, $5) RETURNING *;
