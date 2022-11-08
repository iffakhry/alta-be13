-- input data
INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr1", "Andi", "andi@mail.com","qwerty", "0812345", "jakarta");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr2", "Andi", "andi@mail.com","qwerty", "0812345", "jakarta");

INSERT INTO users(id, nama, email, password, phone, domisili)
VALUES ("usr3", "Budi", "budi@mail.com","qwerty", "0813456", "surabaya");

-- domisili null
INSERT INTO users(id, nama, email, password, phone)
VALUES ("usr4", "Budi", "budi@mail.com","qwerty", "0813456");

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