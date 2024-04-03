UPDATE todo
    LEFT JOIN sharedTodo
    ON todo.id = sharedTodo.todoID
SET isCompleted = ?
WHERE todo.id = ? AND ? IN (todo.userID, sharedTodo.userID);