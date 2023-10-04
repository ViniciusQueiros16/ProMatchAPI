-- DROP TABLE IF EXISTS auth_tokens;
-- DROP TABLE IF EXISTS profile;
-- DROP TABLE IF EXISTS posts;
-- DROP TABLE IF EXISTS matches;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS user_types;

-- Create user_types table
CREATE TABLE user_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    type_name VARCHAR(100) UNIQUE
);

-- Insert user types
INSERT INTO user_types (type_name) VALUES
    ('client'),
    ('professional');

-- Create users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL,
    user_type_id INT NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

-- Insert user data
INSERT INTO users (username, name, email, password, user_type_id, verified) VALUES
    ('silsil', 'Silvio', 'silvio@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 1, TRUE), 
    ('Tony', 'Antônio', 'tony@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 1, TRUE), 
    ('Lu', 'Luciana', 'luciana@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 2, FALSE),  
    ('Sami', 'Samara', 'samara@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 2, FALSE);

-- Create auth_tokens table
CREATE TABLE auth_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(100) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create profile table
CREATE TABLE profile (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    avatar VARCHAR(255),
    birthdate DATE,
    company VARCHAR(100),
    gender VARCHAR(10),
    about TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert profile data
INSERT INTO profile (user_id, avatar, birthdate, company, gender, about, created_at)
VALUES
    (1, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1990-06-15', 'ABC Company', 'Male', 'About Silvio...', NOW()),
    (2, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1985-12-25', 'XYZ Corporation', 'Female', 'About Antônio...', NOW()),
    (3, 'https://avatars.githubusercontent.com/ViniciusQueiros16', '1998-03-10', 'Sample Corp', 'Other', 'About Luciana...', NOW());

-- Create posts table
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    message TEXT,
    image_url VARCHAR(255),
    communityType VARCHAR(50),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert post data
INSERT INTO posts (user_id, message, image_url, communityType)
VALUES
    (1, 'Primeiro post! #início',
     'https://img.freepik.com/fotos-gratis/respingo-colorido-abstrato-3d-background-generativo-ai-background_60438-2509.jpg?w=1480&t=st=1692138193~exp=1692138793~hmac=ada296c954bf989dad8a4f484b363e0a73b9d8a54fa0e2ff87cc69393025e1c3', 'AnyOne'),
    (2, 'Olá mundo! #saudações', 'https://media.istockphoto.com/id/459369173/pt/foto/linda-borboleta-isolado-a-branco.jpg?s=2048x2048&w=is&k=20&c=1WDZr2jioNvgg8ll3CwuCJAKfhZiKUM8W2YNsEF78YQ=', 'Group'),
    (3, 'Compartilhando uma foto incrível.', 'https://www.istockphoto.com/br/foto/nas-asas-da-liberdade-p%C3%A1ssaros-que-voam-e-correntes-quebradas-conceito-da-carga-gm1141549703-305881902', 'Twitter'),
    (4, 'Explorando lugares novos.', 'https://avatars.githubusercontent.com/ViniciusQueiros16', 'AnyOne');

-- Create matches table
CREATE TABLE matches (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    matched_user_id INT NOT NULL,
    is_accepted BOOLEAN DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE KEY unique_match (user_id, matched_user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (matched_user_id) REFERENCES users(id) ON DELETE CASCADE
);
