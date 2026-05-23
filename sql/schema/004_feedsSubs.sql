-- +goose Up
CREATE TABLE feedSubs (
    id UUID NOT NULL UNIQUE, 
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP not Null,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE cascade,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON delete cascade, 
    unique(user_id, feed_id) 
);
-- +goose Down
DROP TABLE feedSubs