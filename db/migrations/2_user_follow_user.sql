-- +goose Up
CREATE TABLE IF NOT EXISTS user_follow_user(
    user_id        INT(10) UNSIGNED NOT NULL,
    follow_user_id INT(10) UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (follow_user_id) REFERENCES user(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
