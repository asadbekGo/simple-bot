drop table if exists media_audio;
drop table if exists media_photo;
drop table if exists media_video;
drop table if exists media_document;

create table media_audio (
    id serial not null primary key,
    from_id numeric not null,
    first_name varchar(120),
    username text not null,
    file_name text not null,
    mime_type text not null,
    file_id text not null,
    file_unique_id text not null,
    file_size numeric,
    duration numeric,
    path text not null,
    created_at timestamp default current_timestamp
);

create table media_photo (
    id serial not null primary key,
    from_id numeric not null,
    first_name varchar(120),
    username text not null,
    file_id text not null,
    file_unique_id text not null,
    file_size numeric,
    width numeric,
    height numeric,
    path text not null,
    created_at timestamp default current_timestamp
);

create table media_video (
    id serial not null primary key,
    from_id numeric not null,
    first_name varchar(120),
    username text not null,
    file_name text not null,
    mime_type text not null,
    file_id text not null,
    file_unique_id text not null,
    file_size numeric,
    width numeric,
    height numeric,
    duration numeric,
    path text not null,
    created_at timestamp default current_timestamp
);

create table media_document (
    id serial not null primary key,
    from_id numeric not null,
    first_name varchar(120),
    username text not null,
    file_name text not null,
    mime_type text not null,
    file_id text not null,
    file_unique_id text not null,
    file_size numeric,
    path text not null,
    created_at timestamp default current_timestamp
);

select file_id from media_photo Where from_id = 700785715;