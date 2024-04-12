CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists products (
    id uuid not null default uuid_generate_v4() unique,
    name varchar(255) not null,
    description text,
    price bigint not null check ( price >= 0 )
);

CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');

create table if not exists orders (
    id uuid not null default uuid_generate_v4() unique,
    order_status status_enum not null default 'created',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);