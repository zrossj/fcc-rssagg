-- +goose Up    
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at timestamp not null,
    modified_at timestamp not null,
    name text not null
);

-- +goose Down
DROP TABLE users; 

