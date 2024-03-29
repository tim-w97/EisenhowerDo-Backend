DROP TABLE IF EXISTS sharedTodo;

CREATE TABLE sharedTodo
(
    id          INT auto_increment NOT NULL,
    todoID      INT                NOT NULL,
    otherUserID INT                NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO sharedTodo (todoID, otherUserID)
VALUES (1, 1),
       (1, 2),
       (2, 3);