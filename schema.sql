create table if not exists users (
    id serial,
    email varchar(255) unique,
    password varchar(60)
);
