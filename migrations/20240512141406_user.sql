-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE `user` (
    `id` BINARY(16) NOT NULL,
    `password` VARCHAR(255) COMMENT 'Пароль',
    `username` VARCHAR(255) UNIQUE COMMENT 'Ник пользователя',
    `first_name` VARCHAR(255) COMMENT 'Имя пользователя',
    `last_name` VARCHAR(255) COMMENT 'Фамилия пользователя',
    `email` VARCHAR(255) UNIQUE COMMENT 'Email пользователя',
    `bio` VARCHAR(255) COMMENT 'Биография пользователя',
    `role` TINYINT COMMENT 'Роль пользователя (0 - default, 10 - admin, 20 - moder)',
    `active` BOOL COMMENT 'Статус активизации аккаунта',
    `verification_code` VARCHAR(20) COMMENT 'Код верификации',
    `verified` BOOL COMMENT 'Статус верификации',
    `created_at` DATETIME COMMENT 'Дата и время создания записи',
    `updated_at` DATETIME COMMENT 'Дата и время последнего обновления записи',
    `deleted_at` DATETIME COMMENT 'Дата и время удаления записи',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE user;