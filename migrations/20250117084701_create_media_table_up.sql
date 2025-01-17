-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS media;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE media (
   media_id INT AUTO_INCREMENT PRIMARY KEY,
   content_id INT NOT NULL,
   file_name VARCHAR(255) NOT NULL,
   file_type VARCHAR(50) NOT NULL,
   file_url TEXT NOT NULL,
   uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP,
   FOREIGN KEY (content_id) REFERENCES content(content_id) ON DELETE CASCADE
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS media;
-- +goose StatementEnd