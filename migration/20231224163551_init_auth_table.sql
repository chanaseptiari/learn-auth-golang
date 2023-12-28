-- +goose Up
-- +goose StatementBegin
CREATE TABLE ms_auth (
 id INT NOT NULL AUTO_INCREMENT,
 Username VARCHAR(255) UNIQUE NOT NULL,
 Password VARCHAR(255),
 Role VARCHAR(100),
 active_at BIGINT,
 created_at BIGINT,
 updated_at BIGINT,
 PRIMARY KEY (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE ms_auth
-- +goose StatementEnd
