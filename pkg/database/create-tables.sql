DROP TABLE IF EXISTS auth_tokens;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO users (username, name, email, password) VALUES
    ('silsil','Silvio', 'silvio@example.com', 'senha1'),
    ('Tony','Antônio', 'tony@example.com', 'senha2'),
    ('Lu','Luciana', 'luciana@example.com', 'senha3'),
    ('Sami','Samara', 'samara@example.com', 'senha4');



CREATE TABLE auth_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(100) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
