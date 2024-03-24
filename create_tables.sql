-- Todo items

DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id    INT auto_increment NOT NULL,
    title VARCHAR(255) NOT NULL,
    text  VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO todo
    (title, text)
VALUES ('Einkaufen', 'Ich brauch noch Toastbrot und Nutella'),
       ('Geschenk für Oma kaufen', 'Ideen: Orchidee, Pralinen, Käsekuchen'),
       ('Bewegen', 'Wenigstens draußen eine Runde um den Block');