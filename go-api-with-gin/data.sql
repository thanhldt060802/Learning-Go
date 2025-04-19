DROP DATABASE IF EXISTS go_api_with_gin;
CREATE DATABASE go_api_with_gin;
USE go_api_with_gin;

CREATE TABLE users (
	user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE invoices (
	invoice_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id),
    total_amount NUMERIC(10, 2),
    status VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(name, email, username, password) VALUES
('Thành', 'admin@gmail.com', 'admid', '1'),
('Phúc', 'user1@gmail.com', 'user1', '1'),
('Tâm', 'user2@gmail.com', 'user2', '1'),
('Hoa', 'user3@gmail.com', 'user3', '1');

INSERT INTO invoices(user_id, total_amount, status) VALUES
(2, 150000, 'Pending'),
(2, 355000, 'Done'),
(2, 250000, 'Done'),
(3, 333000, 'Pending');