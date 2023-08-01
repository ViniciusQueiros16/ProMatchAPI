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


CREATE TABLE profile (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_user INT NOT NULL,
    avatar VARCHAR(255),
    birthdate DATE,
    company VARCHAR(100),
    gender VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (id_user) REFERENCES users (id) ON DELETE CASCADE
);


INSERT INTO profile (id_user, avatar, birthdate, company, gender, created_at)
VALUES
    (1, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1990-06-15', 'ABC Company', 'Male', NOW()),
    (2, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1985-12-25', 'XYZ Corporation', 'Female', NOW()),
    (3, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1998-03-10', 'Sample Corp', 'Other', NOW());

