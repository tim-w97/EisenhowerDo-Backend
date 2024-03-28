DROP TABLE IF EXISTS user;

CREATE TABLE user
(
    id       INT auto_increment NOT NULL,
    username VARCHAR(50) NOT NULL,
    password CHAR(128)   NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO user
    (username, password)
VALUES ('paula', SHA2('ananas', 512)),
       ('ronny', SHA2('benz', 512)),
       ('lisa', SHA2('blume', 512));
