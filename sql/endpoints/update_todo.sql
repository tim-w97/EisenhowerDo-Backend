UPDATE todo,
    (SELECT ? AS newTitle) AS newTitle,
    (SELECT ? AS newText) AS newText,
    (SELECT ? AS newIsImportant) AS newIsImportant,
    (SELECT ? AS newIsUrgent) AS newIsUrgent
SET title      = newTitle,
    text       = newText,
    isImportant = newIsImportant,
    isUrgent = newIsUrgent
WHERE id = ?
  AND userID = ?;
