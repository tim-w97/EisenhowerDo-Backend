UPDATE todo
SET title      = ?,
    text       = ?,
    categoryID = ?
WHERE id = ?
  AND userID = ?;