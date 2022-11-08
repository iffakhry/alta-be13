CREATE DATABASE db_be13;

DROP DATABASE db_be13_test;

USE db_be13;

-- membuat table user
create table users (
id int primary key,
nama varchar(255) not null,
email varchar(100) not null,
phone varchar(15)
);

-- memodifikasi table, menambah kolom password
ALTER TABLE users
ADD password varchar(255);