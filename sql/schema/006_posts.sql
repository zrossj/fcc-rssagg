-- +goose up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title text NOT NULL,
    description text,
    published_at TIMESTAMP NOT NULL,
    url text NOT NULL UNIQUE,
    feed_id UUID NOT NULL references users(id) ON DELETE CASCADE
);

-- +goose down
DELETE TABLE posts; 