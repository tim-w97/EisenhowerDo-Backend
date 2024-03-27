-- Users

DROP TABLE IF EXISTS user;

CREATE TABLE user
(
    id       INT auto_increment NOT NULL,
    username VARCHAR(50),
    password CHAR(128),

    PRIMARY KEY (id)
);

INSERT INTO user
    (username, password)
VALUES ('paula',
        'b6cd1f8dd662768a9f87348f7e88db0acedc913056430f9276855685518e9a7cd3da3db849785687df3599299960a2db986d124de5a54cdcc7fc88c0cde2c4c1'),
       ('ronny',
        '885c5fab5ae64a402aa15b0cb4867ab707d6d76bcb779f6a91b943fda213a12890bef0d882b5efa4c7978127d40a3e89125cda75c961797563dd2e0608db1ea4'),
       ('lisa',
        '6b38f130b7720334b39d1e68cf75c0b4d81fb7267d5ae46af0e0018a031e9bf5861e20a660c3895cc54d2c800e3052270172539817ac6d730b9f333ff7758f0b');


-- Todo items

DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id    INT auto_increment NOT NULL,
    title VARCHAR(100) NOT NULL,
    text  VARCHAR(200) NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO todo
    (title, text)
VALUES ('Einkaufen', 'Ich brauch noch Toastbrot und Nutella'),
       ('Geschenk für Oma kaufen', 'Ideen: Orchidee, Pralinen, Käsekuchen'),
       ('Bewegen', 'Wenigstens draußen eine Runde um den Block');