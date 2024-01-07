-- +migrate Up

CREATE TABLE permissions
(
    id    int primary key AUTO_INCREMENT,
    title varchar(191) not null UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE permissions;