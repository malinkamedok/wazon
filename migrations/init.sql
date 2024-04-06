CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists products (
    id uuid not null default uuid_generate_v4(),
    name varchar(255) not null,
    description text,
    price bigint not null check ( price >= 0 )
);
