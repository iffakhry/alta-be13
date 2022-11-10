-- input data
INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr1", "Andi", "andi@mail.com","qwerty", "0812345", "jakarta");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr2", "Andi", "andi1@mail.com","qwerty", "0812378", "jakarta");

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


-- INNER JOIN
SELECT books.id, books.title, books.author, books.user_id, users.nama, users.email FROM books 
INNER JOIN users ON books.user_id = users.id;

SELECT b.id, b.title, b.author, b.user_id, u.nama, u.email FROM books b
INNER JOIN users u ON b.user_id = u.id;

SELECT books.id, books.title, books.author, books.user_id, users.nama, users.email FROM users 
INNER JOIN books ON books.user_id = users.id
WHERE books.title like "%pemrograman%" ORDER BY books.id desc;

-- LEFT JOIN
SELECT users.id, users.nama, users.email, books.id as book_id, books.title FROM users
LEFT JOIN books ON users.id = books.user_id;

-- RIGHT JOIN
SELECT users.id, users.nama, users.email, books.id, books.title FROM users
RIGHT JOIN books ON users.id = books.user_id;

-- Insert table loan
insert into loan(user_id, book_id, return_date, status)
values ("usr2", 1, "2022-11-12 10:00:00", "dipinjam");

insert into loan(user_id, book_id, return_date, status)
values ("usr5", 8, "2022-11-20", "dipinjam");

insert into loan(user_id, book_id, return_date, status)
values ("usr4", 10, "2022-11-20", "dipinjam");

insert into loan(user_id, book_id, return_date, status)
values ("usr5", 10, "2022-11-25", "dipinjam");

select * from loan;

-- AGREGASI
-- MIN
select min(id) from loan;
select min(return_date) from loan;

-- max
select max(id) from loan;
select max(return_date) as tanggal_kembali from loan;

-- COUNT
select count(id) as jumlah_buku from books;
select count(id) as jumlah_buku from books where user_id = "usr4";

-- having
-- menampilkan user yang memiliki buku, beserta jumlah buku yang dipunyai
select users.id, users.nama, users.email, count(books.id) as jumlah_buku from users
inner join books on users.id = books.user_id
GROUP BY users.id;

-- menampilkan user yang memiliki buku, beserta jumlah buku yang dipunyai, dan hanya menampilkan user yang memiliki buku lebih dari 1
select users.id, users.nama, users.email, count(books.id) as jumlah_buku from users
inner join books on users.id = books.user_id
GROUP BY users.id
HAVING count(books.id) > 1;

SELECT * FROM users;
-- SUBQUERY
-- tampilkan data user yang sudah memiliki buku saja
SELECT user_id FROM books GROUP BY user_id;
SELECT DISTINCT user_id FROM books;

SELECT * FROM users WHERE id IN (SELECT DISTINCT user_id FROM books);

-- tampilkan data user yang belum memiliki buku saja
SELECT * FROM users WHERE id NOT IN (SELECT DISTINCT user_id FROM books);

-- tampilkan user yang belum pernah meminjam buku
select * from loan;
select * from users where id not in (select distinct user_id from loan);

-- tampilkan semua buku yang sudah pernah dipinjam
select * from books where id in (select distinct book_id from loan);

-- tampilkan semua buku yang sudah pernah dipinjam lebih dari 1x
select * from books where id in (select book_id from loan GROUP BY book_id having count(book_id)>1);

select loan.book_id,sum(book_id) from loan group by book_id;


select * from books;
-- MEMANGGIL FUNCTION
SELECT f_count_book_peruser("usr2") as jumlah_buku;

select count(*) from books where user_id = "usr2";

select * from users;
select * from books;
-- tmpilkan data peminjaman buku
SELECT loan.id, loan.user_id, users.nama as nama_peminjam, loan.book_id, books.title, books.user_id as id_pemilik, u.nama as pemilik, loan.return_date, loan.status 
FROM loan
INNER JOIN users ON loan.user_id = users.id
INNER JOIN books ON loan.book_id = books.id
INNER JOIN users u ON books.user_id = u.id;












