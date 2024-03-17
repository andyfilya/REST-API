CREATE TABLE users (
  user_id serial not null unique,
  username varchar(255) not null,
  password varchar(255) not null,
  create_time TIMESTAMP not null,
  user_role varchar(50) not null
);

CREATE TABLE actors (
  actor_id serial not null unique,
  actor_name varchar(255) not null,
  actor_surname varchar(255) not null,
  actor_birth_date date
);

CREATE TABLE films (
  film_id serial not null unique,
  film_title varchar(150) not null,
  film_description varchar(1000) not null,
  film_date date,
  film_rate real
);

CREATE TABLE actors_films (
  id      serial                                           not null unique,
  a_id int references actors (actor_id) on delete cascade      not null,
  f_id int references films (film_id) on delete cascade not null
);

INSERT INTO users (user_id, username, password, create_time, user_role)
VALUES (1, 'ADMIN', '$2a$10$2xtX6brWYGpBFwDU.juBbu2tOx/2MNZDdoBn46ETtlCCQAIfT3TSK', NOW(), 'admin');