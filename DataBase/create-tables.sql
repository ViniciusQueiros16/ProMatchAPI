DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name, email, password) VALUES
    ('Silvio', 'silvio@example.com', 'senha1'),
    ('Tony', 'tony@example.com', 'senha2'),
    ('Luciana', 'luciana@example.com', 'senha3'),
    ('Samara', 'samara@example.com', 'senha4');
