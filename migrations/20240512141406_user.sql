-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE user (
    id binary(16) NOT NULL,
    password VARCHAR(255) NOT NULL COMMENT 'Пароль',
    username VARCHAR(255) UNIQUE NOT NULL COMMENT 'Ник пользователя',
    first_name VARCHAR(255) DEFAULT NULL COMMENT 'Имя пользователя',
    last_name VARCHAR(255) DEFAULT NULL COMMENT 'Фамилия пользователя',
    email VARCHAR(255) UNIQUE NOT NULL COMMENT 'Email пользователя',
    phone VARCHAR(20) UNIQUE DEFAULT NULL COMMENT 'Номер телефона',
    bio VARCHAR(255) DEFAULT NULL COMMENT 'Биография пользователя',
    role VARCHAR(255) NOT NULL COMMENT 'Роль пользователя',
    active BOOL DEFAULT FALSE COMMENT 'Статус активизации аккаунта',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at datetime DEFAULT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE user;
