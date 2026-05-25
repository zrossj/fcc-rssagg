-- name: InsertPosts :one
INSERT INTO posts (
    id, created_at, updated_at, 
    title, description, published_at, url, feed_id)
VALUES (
    $1, $2, $3, 
    $4, $5, $6, $7, $8
)
RETURNING *;


-- name: GetLastPosts :many
SELECT p.* FROM posts p
JOIN feedsubs fs on fs.feed_id = p.feed_id 
JOIN users u ON u.id = fs.user_id 
WHERE u.id = $1 
LIMIT $2;