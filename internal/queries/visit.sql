-- name: CreateVisit :one
INSERT INTO visit (short_url_id, user_agent, occurred_at)
SELECT s.id, $2, $3
FROM short_url as s
WHERE s.slug = $1
RETURNING *;
