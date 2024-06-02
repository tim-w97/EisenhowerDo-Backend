SELECT todo.id, todo.userID, title, text, isCompleted
FROM todo
    LEFT JOIN sharedTodo
    ON todo.id = sharedTodo.todoID
WHERE todo.id = ? AND ? IN (todo.userID, sharedTodo.userID);