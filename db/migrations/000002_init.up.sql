create schema if not exists schema_name;

create table if not exists schema_name.urls
(
    id         uuid    not null
        constraint urls_pk
            primary key
        default gen_random_uuid(),
    short_id   varchar(10) not null
        constraint urls_short_id_uq
            unique,
    url        text    not null,
    created_at timestamp default current_timestamp,
    expire_at  timestamp default (current_timestamp + interval '1 day')
);
