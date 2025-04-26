create schema if not exists schema_name;

create table if not exists schema_name.statistic
(
    id           uuid default gen_random_uuid() primary key,
    short_id     varchar(10) not null,
    url          text not null,
    clicked      timestamp default current_timestamp,
    ip_address   varchar(45),
    user_agent   varchar(512),
    country      varchar(100),
    device_type  varchar(100)
);