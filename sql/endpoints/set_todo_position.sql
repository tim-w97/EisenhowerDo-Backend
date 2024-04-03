UPDATE todo,
    (SELECT id, position FROM todo WHERE id = ?) AS todoToMove,
    (SELECT ? AS desiredPosition) AS desiredPosition
SET todo.position =
    CASE
        WHEN todo.position < todoToMove.position AND todo.position >= desiredPosition THEN todo.position + 1
        WHEN todo.position > todoToMove.position AND todo.position <= desiredPosition THEN todo.position - 1
        WHEN todo.id = todoToMove.id THEN desiredPosition
        ELSE todo.position
    END
WHERE userID = ?;