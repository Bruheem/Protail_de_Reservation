CREATE TRIGGER update_library_subscriber_count
AFTER INSERT ON subscription
FOR EACH ROW
BEGIN
  UPDATE library
  SET NumSubscribers = NumSubscribers + 1
  WHERE ID = NEW.library_id;
END$$

CREATE TRIGGER decrement_library_subscriber_count
AFTER DELETE ON subscription
FOR EACH ROW
BEGIN
  UPDATE library
  SET NumSubscribers = NumSubscribers - 1
  WHERE ID = OLD.library_id;
END$$

CREATE TRIGGER prevent_double_borrowing
BEFORE INSERT ON lending
FOR EACH ROW
BEGIN
  IF EXISTS (
    SELECT 1
    FROM lending
    WHERE user_id = NEW.user_id
      AND document_id = NEW.document_id
      AND return_date IS NULL
  ) THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'User has already borrowed this document.';
  END IF;
END$$