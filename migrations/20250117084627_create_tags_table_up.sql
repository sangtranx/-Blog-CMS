-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tags (
  tag_id INT AUTO_INCREMENT PRIMARY KEY,
  tag_name VARCHAR(100) NOT NULL UNIQUE,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tags;
-- +goose StatementEnd