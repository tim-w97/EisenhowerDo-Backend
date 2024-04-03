SELECT todo.id, userID, title, text, categoryID, position, isCompleted
FROM sharedTodo
INNER JOIN todo
ON todo.id = sharedTodo.todoID
WHERE sharedTodo.otherUserID = ?;