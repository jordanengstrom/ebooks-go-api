

CREATE TABLE books (
    ID UUID DEFAULT uuid_generate_v4(),
    author_id UUID REFERENCES authors(ID),
    title VARCHAR(50),
    copyright_year INT,
    about TEXT,
    PRIMARY KEY(ID)
);

INSERT INTO books (
    author_id,
    title,
    copyright_year,
    about
) VALUES 
    ('827d1b7c-c9a0-49fa-b0ba-3e62d5b9ec1e', 'This Side of Paradise', 1920, 'The Jazz Age was not just a drastic change in the cluture; it was a new dimension in consciousness. During the 1910s and 1920s America underwent a massive paradigm shift, a transition from an era of smug Victorian conformity and certainty to one of confusion and ambiguity called "modernism".'),
    ('ce4793ac-0b7a-4e71-a593-7f1ca0e23bf0', 'The Alchemist', 1988, 'Paulo Coelho''s enchanting novel has inspired a devoted following around the world. This story, dazzling in its powerful simplicity and inspiring wisdom, is about an Andalusian shepherd boy named Santiago who travels from his homeland in Spain to the Egyptian desert in search of a treasure buried in the Pyramids.'),
    ('7a86cbd0-02c2-41ed-a52a-9ebddcbaab10', 'Dracula', 1897, 'Written in the form of letters, diary entries, and news bits, Dracula chronicles the vampire''s journey from his Transylvanian castle to the nighttime streets of London. There he searches for the blood he needs to stay alive - the blood of strong men and beautiful women - while his enemies plot to rid the world of his frightful power.');


SELECT * FROM books;

SELECT * FROM information_schema.tables
WHERE table_schema NOT IN ('information_schema', 'pg_catalog');

DROP TABLE books;

CREATE TABLE authors (
    ID UUID DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    PRIMARY KEY(ID)
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