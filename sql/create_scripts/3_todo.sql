DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id          INT auto_increment NOT NULL,
    userID      INT                NOT NULL,
    title       VARCHAR(100)       NOT NULL,
    text        VARCHAR(200)       NOT NULL,
    categoryID  INT                NOT NULL,
    position    INT                NOT NULL,
    isCompleted BOOL DEFAULT false NOT NULL,

    PRIMARY KEY (id),

    FOREIGN KEY (userID) REFERENCES user (id),
    FOREIGN KEY (categoryID) REFERENCES category (id)
);

INSERT INTO todo (userID, title, text, categoryID, position)
VALUES (1, 'Einkaufen', 'Ich brauch noch Toastbrot und Nutella', 1, 1),
       (1, 'Geschenk für Oma kaufen', 'Ideen: Orchidee, Pralinen, Käsekuchen', 2, 2),
       (1, 'Bewegen', 'Wenigstens draußen eine Runde um den Block', 2, 3);