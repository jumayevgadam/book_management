CREATE TABLE IF NOT EXISTS authors (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   biography TEXT,
   birthdate DATE
);

CREATE TABLE IF NOT EXISTS books (
   id SERIAL PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   author_id INT REFERENCES authors(id) ON DELETE CASCADE NOT NULL,
   year INT,
   genre VARCHAR(255)
);  