-- +migrate Up

alter table users
    ADD COLUMN role ENUM ('mysqluser','admin') not null;


-- +migrate Down
alter table users
    drop COLUMN role;