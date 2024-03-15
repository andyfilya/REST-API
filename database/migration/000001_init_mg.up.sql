CREATE TABLE users (
  user_id serial not null unique,
  username varchar(255) not null,
  password varchar(255) not null,
  create_time TIMESTAMP not null
);
