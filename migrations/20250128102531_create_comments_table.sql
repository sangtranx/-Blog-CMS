-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE comments (
      comment_id INT AUTO_INCREMENT PRIMARY KEY,
      content_id INT NOT NULL,
      user_id INT NOT NULL,
      comment_text TEXT NOT NULL,
      comment_date DATETIME DEFAULT CURRENT_TIMESTAMP,
      FOREIGN KEY (content_id) REFERENCES content(content_id) ON DELETE CASCADE,
      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
