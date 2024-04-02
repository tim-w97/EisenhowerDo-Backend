SELECT *
FROM todo
WHERE userID = ?
   OR id IN
      (SELECT todoID
       FROM sharedTodo
       WHERE otherUserID = ?);