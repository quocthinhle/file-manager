-- name: GetAllNodes :many
SELECT * FROM node;

-- name: GetParentNodes :many
SELECT * FROM node WHERE parent_id IS NULL AND owner_id = $1;