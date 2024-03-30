UPDATE todo
SET isCompleted = ?
WHERE id = ?
  AND userID = ?;