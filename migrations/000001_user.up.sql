CREATE TABLE IF NOT EXISTS users(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(32),
    password BINARY(60),
    name VARCHAR(64),
    UNIQUE (username)
);
