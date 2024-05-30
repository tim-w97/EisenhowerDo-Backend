UPDATE todo,
    (SELECT ? AS newTitle) AS newTitle,
    (SELECT ? AS newText) AS newText
SET title      = IF(newTitle = "", title, newTitle),
    text       = IF(newText = "", text, newText)
WHERE id = ?
  AND userID = ?;
