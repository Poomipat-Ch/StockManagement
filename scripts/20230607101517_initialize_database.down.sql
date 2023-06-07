create table if not exists users (
    id serial primary key,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(194) not null unique,
    password_hashed varchar(100) not null,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp
);