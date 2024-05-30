SELECT todo.id, todo.userID, title, text, position, isCompleted
FROM sharedTodo
    INNER JOIN todo
    ON todo.id = sharedTodo.todoID
WHERE sharedTodo.userID = ?;