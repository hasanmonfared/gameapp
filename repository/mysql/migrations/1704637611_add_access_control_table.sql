-- +migrate Up

CREATE TABLE access_control
(
    id            int primary key AUTO_INCREMENT,
    actor_id      int                       not null,
    actor_type    ENUM ('mysqluser','role') NOT NULL,
    permission_id INT                       NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (permission_id) REFERENCES permissions (id)
);

-- +migrate Down
DROP TABLE access_control;