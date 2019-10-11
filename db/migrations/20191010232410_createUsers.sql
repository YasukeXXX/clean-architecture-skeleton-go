
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id serial not null,
  name varchar(127) not null,
  email varchar(127) not null,
  password_digest varchar(255) not null,
  created_at timestamp not null,
  updated_at timestamp not null,
  primary key (id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;

