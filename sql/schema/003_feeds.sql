-- +goose Up
CREATE TABLE feeds (
    id UUID NOT NULL UNIQUE, 
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP not Null,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id)  ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds