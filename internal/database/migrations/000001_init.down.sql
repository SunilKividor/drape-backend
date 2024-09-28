-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS auth_method;