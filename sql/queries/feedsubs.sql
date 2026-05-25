-- name: SubFeed :one
INSERT INTO feedSubs (id, created_at, modified_at, user_id, feed_id) 
    VALUES ($1, $2, $3, $4, $5
    )
RETURNING *;

-- name: GetUserFeeds :many
SELECT * FROM feedsubs f where f.user_id = $1;

-- name: DeleteUserSub :exec
DELETE FROM feedsubs f 
where f.user_id = $1 and f.feed_id = $2; 

