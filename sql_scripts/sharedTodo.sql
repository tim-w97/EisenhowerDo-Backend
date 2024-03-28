DROP TABLE IF EXISTS sharedTodo;

CREATE TABLE sharedTodo
(
    id               INT auto_increment NOT NULL,
    todoID           INT                NOT NULL,
    authorizedUserID INT                NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO sharedTodo (todoID, authorizedUserID)
VALUES (1, 1),
       (1, 2),
       (2, 3);