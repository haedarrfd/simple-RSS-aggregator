-- +goose Up
CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  last_fetched_at TIMESTAMP,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;