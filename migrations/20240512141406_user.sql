-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE user (
    id binary(16) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL COMMENT 'Email пользователя',
    phone VARCHAR(20) UNIQUE NOT NULL COMMENT 'Номер телефона',
    password VARCHAR(255) NOT NULL COMMENT 'Пароль',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
