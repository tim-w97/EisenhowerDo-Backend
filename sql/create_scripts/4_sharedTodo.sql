DROP TABLE IF EXISTS sharedTodo;

CREATE TABLE sharedTodo
(
    todoID INT NOT NULL,
    userID INT NOT NULL,

    PRIMARY KEY (todoID, userID),

    FOREIGN KEY (todoID) REFERENCES todo (id),
    FOREIGN KEY (userID) REFERENCES user (id)
);

INSERT INTO sharedTodo (todoID, userID)
VALUES (1, 2),
       (3, 2);