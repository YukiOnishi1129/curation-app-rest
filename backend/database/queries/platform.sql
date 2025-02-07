-- name: GetPlatform :one
SELECT * FROM platforms
WHERE id = $1 LIMIT 1;

-- name: ListPlatforms :many
SELECT * FROM platforms
ORDER BY name;

-- name: CreatePlatforms :one
INSERT INTO platforms (
  name, site_url, platform_site_type, favicon_url
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdatePlatforms :exec
UPDATE platforms
  set name = $2,
  site_url = $3,
  platform_site_type = $4,
  favicon_url = $5
WHERE id = $1;

-- name: DeletePlatforms :exec
DELETE FROM platforms
WHERE id = $1;