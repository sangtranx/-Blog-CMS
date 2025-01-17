-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS content;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE content (
     content_id INT AUTO_INCREMENT PRIMARY KEY,
     content_type ENUM('article', 'blog', 'page') NOT NULL,
     title VARCHAR(255) NOT NULL,
     content_description TEXT NOT NULL,
     user_id INT NOT NULL,
     category_id INT,
     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
     FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS content;
-- +goose StatementEnd