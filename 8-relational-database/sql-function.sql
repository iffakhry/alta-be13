DELIMITER $$
CREATE FUNCTION f_count_book_peruser (user_id_p CHAR(5)) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM books WHERE user_id = user_id_p;
RETURN total;
END$$
DELIMITER ;