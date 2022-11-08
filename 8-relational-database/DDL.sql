CREATE DATABASE db_be13;

DROP DATABASE db_be13_test;

USE db_be13;

-- memodifikasi table, menambah kolom password
ALTER TABLE users
ADD password varchar(255);

-- menghapus table
DROP TABLE users;

-- membuat table user
create table users (
id varchar(10) not null primary key,
nama varchar(255) not null,
email varchar(100) not null,
password varchar(255) not null,
phone varchar(15),
domisili varchar(100)
);

CREATE table books(
id int not null primary key AUTO_INCREMENT,
title varchar(255) not null,
publisher varchar(255),
author varchar(100),
publish_year varchar(50),
user_id varchar(10),
CONSTRAINT fk_BooksUsers FOREIGN KEY (user_id) REFERENCES users(id)
);


create table loan(
id int not null primary key AUTO_INCREMENT,
user_id varchar(10),
book_id int,
return_date datetime,
status ENUM("dipinjam", "dikembalikan"),
created_at datetime default current_timestamp,
updated_at datetime default current_timestamp,
CONSTRAINT fk_LoanUsers FOREIGN KEY (user_id) REFERENCES users(id),
CONSTRAINT fk_LoanBooks FOREIGN KEY (book_id) REFERENCES books(id)
);









