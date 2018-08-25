create table if not exists users (
    id bigserial primary key,
    email varchar(254) unique,
    password varchar(60)
);

create table if not exists sessions (
    id bigserial primary key,
    user_id bigint references users (id) on delete cascade,
    token varchar(64)
);
