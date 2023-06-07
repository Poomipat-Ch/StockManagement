create table if not exists users (
    id integer primary key autoincrement,
    first_name text min 2 max 100 not null,
    last_name text min 2 max 100 not null,
    email text max 194 not null unique,
    password_hashed text min 8 max 100 not null,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp
)