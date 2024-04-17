CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE SCHEMA IF NOT EXISTS productcard;
SET search_path TO productcard;
create table if not exists products (
    id uuid not null default public.uuid_generate_v4(),
    name varchar(255) not null,
    description text,
    price bigint not null check ( price >= 0 )
    );

CREATE SCHEMA IF NOT EXISTS accountservice;
SET search_path TO accountservice;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE user_account (
                              id uuid default public.uuid_generate_v4() PRIMARY KEY,
                              name varchar(100),
                              surname varchar(100),
                              login varchar(100),
                              password varchar(100),
                              FOREIGN KEY (id) REFERENCES user_account (id)
);

CREATE TABLE cart (
                      id uuid default public.uuid_generate_v4() PRIMARY KEY,
                      userID uuid REFERENCES user_account (id)
);

CREATE TABLE product (
                         id uuid PRIMARY KEY,
                         name varchar(100),
                         description varchar(255),
                         price int
);
CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');
CREATE TABLE u_order (
                           id uuid PRIMARY KEY,
                           userID uuid REFERENCES user_account (id),
                           order_status status_enum not null,
                           createdAt timestamp with time zone DEFAULT current_timestamp,
                           updatedAt timestamp with time zone DEFAULT current_timestamp
);


CREATE TABLE product_order (
                               productID uuid REFERENCES product (id),
                               orderID uuid REFERENCES u_order (id),
                               PRIMARY KEY (productID, orderID)
);

CREATE TABLE product_cart (
                             cartID uuid REFERENCES cart (id),
                             productID uuid REFERENCES product (id),
                             PRIMARY KEY (cartID, productID)
);

CREATE SCHEMA IF NOT EXISTS order_service;
SET search_path TO order_service;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE status_enum AS ENUM ('created', 'prepare', 'delivery', 'await', 'received');

create table if not exists orders (
    id uuid default public.uuid_generate_v4() PRIMARY KEY,
    order_status status_enum not null default 'created',
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
