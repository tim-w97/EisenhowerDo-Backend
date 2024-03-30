DROP TABLE IF EXISTS category;

CREATE TABLE category
(
    id       INT auto_increment NOT NULL,
    category VARCHAR(100)       NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO category (category)
VALUES ('Privat'),
       ('Schule'),
       ('Einkaufen');