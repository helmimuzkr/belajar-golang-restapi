CREATE TABLE IF NOT EXISTS categories (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(25) NOT NULL
);

DESC categories;

INSERT INTO categories(name)
    VALUES('hardware');

SELECT * FROM categories;