CREATE TABLE users (
  user_id serial not null unique,
  username varchar(255) not null,
  password varchar(255) not null,
  create_time TIMESTAMP not null
);

CREATE TABLE actors (
  actor_id serial not null unique,
  actor_name varchar(255) not null,
  actor_surname varchar(255) not null,
  actor_birth_date date
);
