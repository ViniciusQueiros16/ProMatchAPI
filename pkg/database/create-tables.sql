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
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL,
    user_type_id INT NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    privacy_accepted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_type_id) REFERENCES user_types(id)
);

-- Insert user data
INSERT INTO users (username, email, password, user_type_id, verified, privacy_accepted) VALUES
    ('silsil', 'silvio@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 1, TRUE, TRUE), 
    ('Tony', 'tony@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 1, TRUE, FALSE), 
    ('Lu', 'luciana@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 2, FALSE, TRUE),  
    ('Sami', 'samara@example.com', '$2a$10$8yaLfN1mVcQZcMis149VZesprNumE0ULCaxKvM2P8mXuZ5eWTkzVG', 2, FALSE, FALSE);

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
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    avatar VARCHAR(255),
    cover_photo VARCHAR(255),
    phone_number VARCHAR(20),
    birthdate DATE,
    gender VARCHAR(10),
    about TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert profile data
INSERT INTO profile (user_id, first_name, last_name, avatar, cover_photo, phone_number, birthdate, gender, about, created_at, updated_at, deleted_at)
VALUES
    (1, 'Silvio', 'Queiros', 'https://avatars.githubusercontent.com/ViniciusQueiros16', NULL, NULL, '1990-06-15', 'Male', 'About Silvio...', NOW()),
    (2, 'Antônio', 'Silva', 'https://avatars.githubusercontent.com/ViniciusQueiros16', NULL, NULL, '1985-12-25', 'Female', 'About Antônio...', NOW()),
    (3, 'Luciana', 'Santos', 'https://avatars.githubusercontent.com/ViniciusQueiros16', NULL, NULL, '1998-03-10', 'Other', 'About Luciana...', NOW());

CREATE TABLE user_addresses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    country VARCHAR(100),
    street_address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    postal_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE notifications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    comments BOOLEAN DEFAULT TRUE,
    candidates BOOLEAN DEFAULT TRUE,
    offers BOOLEAN DEFAULT TRUE,
    sms_delivery_option ENUM('push-everything', 'push-email', 'no-push-notifications') DEFAULT 'push-everything',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

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
    (3, 'Compartilhando uma foto incrível.', 'https://www.istockphoto.com/br/foto/nas-asas-da-liberdade-pássaros-que-voam-e-correntes-quebradas-conceito-da-carga-gm1141549703-305881902', 'Twitter'),
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

-- Create professional table
CREATE TABLE professional (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    type_service VARCHAR(100),
    recommendations INT,
    professional_experience TEXT,
    link_Social_media VARCHAR(255),
    link_portfolio VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create client table
CREATE TABLE client (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    company VARCHAR(100),
    recommendations INT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create services table
CREATE TABLE services (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255),
    description TEXT,
    time_to_receive_proposals INT,
    category VARCHAR(100),
    service_visibility ENUM('public', 'private', 'Completed'),
    attach_files VARCHAR(255),
    status ENUM('published', 'open', 'Completed'),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE skills (
    skill_id INT AUTO_INCREMENT PRIMARY KEY,
    skill_name VARCHAR(100) UNIQUE
);

CREATE TABLE professional_skills (
    professional_id INT NOT NULL,
    skill_id INT NOT NULL,
    PRIMARY KEY (professional_id, skill_id),
    FOREIGN KEY (professional_id) REFERENCES professional(user_id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id) ON DELETE CASCADE
);

CREATE TABLE service_desired_skills (
    service_id INT NOT NULL,
    skill_id INT NOT NULL,
    PRIMARY KEY (service_id, skill_id),
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id) ON DELETE CASCADE
);

-- Create ratings table
CREATE TABLE ratings (
    rating_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    rated_by_user_id INT NOT NULL,
    rating DECIMAL(3, 2) NOT NULL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (rated_by_user_id) REFERENCES users(id) ON DELETE CASCADE
);
