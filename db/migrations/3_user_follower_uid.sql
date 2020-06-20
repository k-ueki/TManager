-- +goose Up
CREATE TABLE IF NOT EXISTS user_follower_tid(
    user_id      INT(10) UNSIGNED NOT NULL,
    follower_tid INT(10) UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
