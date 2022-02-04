CREATE TABLE books (
    id SERIAL,
    author_id INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    copyright_year INT,
    about VARCHAR(355),
    PRIMARY KEY(id)
)

INSERT INTO books (
    author_id,
    title,
    copyright_year,
    about
) VALUES 
    (1, "This Side of Paradise", 1920, "The Jazz Age was not just a drastic change in the cluture; it was a new dimension in consciousness. During the 1910s and 1920s America underwent a massive paradigm shift, a transition from an era of smug Victorian conformity and certainty to one of confusion and ambiguity called 'modernism'."),
    (2, "The Alchemist", 1988, "Paulo Coelho's enchanting novel has inspired a devoted following around the world. This story, dazzling in its powerful simplicity and inspiring wisdom, is about an Andalusian shepherd boy named Santiago who travels from his homeland in Spain to the Egyptian desert in search of a treasure buried in the Pyramids."),
    (3, "Dracula", 1897, "Written in the form of letters, diary entries, and news bits, Dracula chronicles the vampire's journey from his Transylvanian castle to the nighttime streets of London. There he searches for the blood he needs to stay alive - the blood of strong men and beautiful women - while his enemies plot to rid the world of his frightful power.");


SELECT * FROM books;

DROP TABLE books;


CREATE TABLE authors (
    id SERIAL,
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    PRIMARY KEY(id)
);

INSERT INTO authors (
    first_name,
    middle_name,
    last_name
) VALUES
    ('F.', 'Scott', 'Fitzgerald'),
    ('Paulo', '', 'Coelho'),
    ('Bram', '', 'Stoker');

SELECT * FROM authors;

DROP TABLE authors;