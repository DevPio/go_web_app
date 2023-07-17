CREATE DATABASE IF NOT EXISTS goweb;

CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(100) NOT NULL,
    publication_year INT,
    genre VARCHAR(50),
    isbn VARCHAR(20) UNIQUE,
    price DECIMAL(10, 2),
    copies_available INT DEFAULT 0
);


INSERT INTO books (title, author, publication_year, genre, isbn, price, copies_available) 
VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction', '978-0743273565', 12.99, 25);

INSERT INTO books (title, author, publication_year, genre, isbn, price, copies_available) 
VALUES ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction', '978-0061120084', 10.50, 30);
