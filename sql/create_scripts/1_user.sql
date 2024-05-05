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
VALUES ('lisa', SHA2('check', 512)),
       ('tim', SHA2('golang', 512));
