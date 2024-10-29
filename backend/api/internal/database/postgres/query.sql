-- name: GetAllNodes :many
SELECT * FROM node;

-- name: GetParentNodes :many
SELECT * FROM node WHERE parent_id IS NULL AND owner_id = $1;

-- name: CreateNode :one
INSERT INTO node (id, type, name, parent_id, owner_id)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetNode :one
SELECT
    n.id as id,
    n.type as type,
    n.name as name,
    n.parent_id as parent_id,
    n.owner_id as owner_id,
    COALESCE(json_agg(json_build_object('id', c.id,
                                       'name', c.name,
                                       'type', c.type,
                                       'parent_id', c.parent_id,
                                       'owner_id', c.owner_id)), '[]') as children
FROM node n
         INNER JOIN node_closure nc on n.id = nc.ancestor_id AND nc.depth = 1
INNER JOIN node c on nc.descendant_id = c.id
WHERE n.id = $1
GROUP BY n.id;