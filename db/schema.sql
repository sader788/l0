create table orders
(
    uuid       varchar(30) not null
        constraint uuid_pk
            primary key,
    order_json jsonb
);