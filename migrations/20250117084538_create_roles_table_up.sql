-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE roles (
   role_id INT AUTO_INCREMENT PRIMARY KEY,
   role_name VARCHAR(50) NOT NULL UNIQUE,
   created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd