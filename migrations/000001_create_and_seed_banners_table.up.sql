CREATE TABLE IF NOT EXISTS banners (
   id BIGINT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   description VARCHAR(255),
   link VARCHAR(255),
   link_type ENUM('URL', 'SCREEN'),
   image_url VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);


INSERT INTO
    banners (title, description, image_url)
VALUES
    ('Your Project Companion', 'by Codespace Indonesia', 'https://minio-cloud.codespace.id/backoffice/rocket.png'),
    ('Codespace X', 'by Codespace Indonesia', 'https://minio-cloud.codespace.id/backoffice/Header.png');