DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id          INT auto_increment NOT NULL,
    userID      INT                NOT NULL,
    title       VARCHAR(100)       NOT NULL,
    text        VARCHAR(200)       NOT NULL,
    categoryID  INT                NOT NULL,
    isImportant BOOL               NOT NULL,
    isUrgent    BOOL               NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (userID) REFERENCES user (id),
    FOREIGN KEY (categoryID) REFERENCES category (id)
);