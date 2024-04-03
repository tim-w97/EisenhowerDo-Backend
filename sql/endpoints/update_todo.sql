UPDATE todo,
    (SELECT ? AS newTitle) AS newTitle,
    (SELECT ? AS newText) AS newText,
    (SELECT ? AS newCategoryID) AS newCategoryID
SET title      = IF(newTitle = "", title, newTitle),
    text       = IF(newText = "", text, newText),
    categoryID = IF(newCategoryID = 0, categoryID, newCategoryID)
WHERE id = ?
  AND userID = ?;
