-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS content_tags;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE content_tags (
  content_id INT NOT NULL,
  tag_id INT NOT NULL,
  PRIMARY KEY (content_id, tag_id),
  FOREIGN KEY (content_id) REFERENCES content(content_id) ON DELETE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tags(tag_id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS content_tags;
-- +goose StatementEnd