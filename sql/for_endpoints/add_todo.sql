INSERT INTO todo (userID, title, text, categoryID, position)
SELECT ? as passedUserID,
       ?,
       ?,
       ?,
       IFNULL(
               (SELECT MAX(position) + 1 FROM todo WHERE userID = passedUserID),
       1);
