UPDATE todo
SET title = ?,
    text  = ?
WHERE id = ?
  AND userID = ?;