create table if not exists users (
    id serial primary key,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(194) not null unique,
    password_hashed varchar(100) not null,
    created_at timestamp not null default (now() at time zone 'utc'),
    updated_at timestamp not null default (now() at time zone 'utc'),
    deleted_at timestamp 
);