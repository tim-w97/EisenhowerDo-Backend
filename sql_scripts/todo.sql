DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id          INT auto_increment NOT NULL,
    userID      INT                NOT NULL,
    title       VARCHAR(100)       NOT NULL,
    text        VARCHAR(200)       NOT NULL,
    position    INT                NOT NULL,
    isCompleted BOOL DEFAULT false NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (userID) REFERENCES user (id)
);

INSERT INTO todo (userID, title, text, position)
VALUES (3, 'Einkaufen', 'Ich brauch noch Toastbrot und Nutella', 1),
       (3, 'Geschenk für Oma kaufen', 'Ideen: Orchidee, Pralinen, Käsekuchen', 2),
       (3, 'Bewegen', 'Wenigstens draußen eine Runde um den Block', 3);