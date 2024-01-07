-- +migrate Up

CREATE TABLE access_control
(
    id            int primary key AUTO_INCREMENT,
    actor_id      varchar(191)         not null UNIQUE,
    actor_type    ENUM ('user','role') NOT NULL,
    permission_id INT                  NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);

-- +migrate Down
DROP TABLE access_control;