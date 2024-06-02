SELECT todo.id, todo.userID, title, text, isImportant, isUrgent
FROM sharedTodo
    INNER JOIN todo
    ON todo.id = sharedTodo.todoID
WHERE sharedTodo.userID = ?;