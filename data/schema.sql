create table if not exists users (
    id bigserial primary key,
    email varchar(254) unique not null,
    password varchar(60) not null,
    created_at timestamp with time zone default current_timestamp
);

alter table users add column if not exists verified boolean not null default false;
alter table users add column if not exists verification_code varchar(64) default null;

create table if not exists sessions (
    id bigserial primary key,
    user_id bigint not null references users (id) on delete cascade,
    token varchar(64) not null,
    created_at timestamp with time zone default current_timestamp
);

create table if not exists notes (
    id bigserial primary key,
    user_id bigint not null references users (id) on delete cascade,
    content text,
    created_at timestamp with time zone default current_timestamp
);
