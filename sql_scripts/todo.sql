DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id          INT auto_increment NOT NULL,
    title       VARCHAR(100)       NOT NULL,
    text        VARCHAR(200)       NOT NULL,
    userID      INT                NOT NULL,
    isCompleted BOOL DEFAULT false NOT NULL,

    PRIMARY KEY (id)
);

INSERT INTO todo (title, text, userID)
VALUES ('Einkaufen', 'Ich brauch noch Toastbrot und Nutella', 3),
       ('Geschenk für Oma kaufen', 'Ideen: Orchidee, Pralinen, Käsekuchen', 3),
       ('Bewegen', 'Wenigstens draußen eine Runde um den Block', 3);