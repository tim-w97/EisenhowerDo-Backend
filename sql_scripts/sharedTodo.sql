DROP TABLE IF EXISTS sharedTodo;

CREATE TABLE sharedTodo
(
    id          INT auto_increment NOT NULL,
    todoID      INT                NOT NULL,
    otherUserID INT                NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (todoID) REFERENCES todo (id),
    FOREIGN KEY (otherUserID) REFERENCES user (id)
);

INSERT INTO sharedTodo (todoID, otherUserID)
VALUES (1, 1),
       (1, 2),
       (2, 3);