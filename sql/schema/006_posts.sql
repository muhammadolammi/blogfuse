-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    description TEXT,
    title TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL UNIQUE,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID NOt NULL REFERENCES feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts;