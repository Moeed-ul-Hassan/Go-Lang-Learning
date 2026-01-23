CREATE DATABASE startersql;
USE startersql;
CREATE TABLE  users {
id INT AUTO_INCREMENT PRIMARY KEY,
 name VARCHAR(100) NOT NULL,
 email VARCHAR(100) NOT NULL,
 gender ENUM('MALE',, 'FEMALE', 'KHUSRAA'),
 date_of_birth DATE,
}
SELECT * FROM users;