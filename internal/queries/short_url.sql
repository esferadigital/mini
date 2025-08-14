-- name: GetShortURL :one
SELECT * FROM short_url
WHERE id = $1 LIMIT 1;

-- name: GetShortURLBySlug :one
SELECT * FROM short_url
WHERE slug = $1 LIMIT 1;

-- name: CreateShortURL :one
INSERT INTO short_url (
  slug, target_url
) VALUES (
  $1, $2
)
RETURNING *;
