-- input data
INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr1", "Andi", "andi@mail.com","qwerty", "0812345", "jakarta");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr2", "Andi", "andi@mail.com","qwerty", "0812345", "jakarta");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr3", "Budi", "budi@mail.com","qwerty", "0813456", "surabaya");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr4", "Anton", "anton@mail.com","qwerty", "0813457", "malang");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr5", "Rudi", "rudi@mail.com","qwerty", "0813487", "malang");

-- domisili null
INSERT INTO users(id, nama, email, password, phone)
VALUES ("usr4", "Anton", "anton@mail.com","qwerty", "0813457");

-- error, nama tidak boleh null
INSERT INTO users(id, email, password, phone)
VALUES ("usr5", "budi@mail.com","qwerty", "0813456");

-- membaca data
SELECT * FROM users;

SELECT id, nama, email, phone from users where id = "usr1";

SELECT id, nama, email, phone from users where nama = "Andi";

SELECT id, nama, email, phone from users where domisili IS NOT NULL;

SELECT id, nama, email, phone from users where domisili IS NULL;


-- update data
UPDATE users SET 
email="andi1@mail.com",
phone="08567"
wHERE id = "usr2";

-- delete data
DELETE FROM users where id = "usr4";

-- like between
SELECT * FROM users where email like "and%";
SELECT * FROM users where email = "andi@mail.com";
SELECT * FROM users where email like "%com";

-- insert books
INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("one piece", "gramedia pustaka", "Oda", "November 2000", "usr1");

INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("dasar pemrograman", "gramedia pustaka", "Oda", "November 2000", "usr1");

INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("boruto", "bentang pustaka", "musashi", "januari 2010", "usr2");

INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("pemrograman golang", "andi publisher", "andi", "januari 2022", "usr3");

INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("pemrograman javascript", "andi publisher", "andi", "januari 2022", "usr3");

INSERT INTO books (title, publisher, author, publish_year, user_id)
VALUES ("pemrograman java", "andi publisher", "andi", "januari 2022", "usr5");

DELETE FROM books where id = 4;
-- read books
select * from users;
SELECT * FROM books;

-- between
select * from books where id between 2 and 6;

-- menampilkan buku yang ada kata "pemrograman" di judulnya
select * from books where title like "%pemrograman%";

select * from books where title like "%pemrograman%" AND publisher = "andi publisher" OR author = "musashi";

-- ORDER BY
select * from books where title like "%pemrograman%" order by id desc;
select * from books where title like "%pemrograman%" order by author asc;

-- LIMIT
select * from books where title like "%pemrograman%" order by id asc limit 2;

-- hapus user 1
delete from users where id = "usr2";
delete from users where id = "usr5";

delete from books where id = 9;

select * from users;
SELECT * FROM books;

-- restrict = harus menghapus di level child dulu, baru bisa menghapus di level parent
-- cascade = bisa langsung menghapus di level parent, namun data di table child akan ikut ke hapus















