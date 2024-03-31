UPDATE todo
SET position = position + 1
WHERE position < ?
  AND position >= ?
  AND userID = ?;