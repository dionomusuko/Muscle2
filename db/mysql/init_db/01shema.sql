CREATE DATABASE IF NOT EXISTS `go-db`;
USE `go-db`;

CREATE TABLE IF NOT EXISTS users(
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `name`  VARCHAR(100) NOT NULL,
    `email` VARCHAR(256) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS tasks(
   `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
   `user_id` BIGINT NOT NULL,
   `menu` VARCHAR(100) NOT NULL,
   `set_count` INT(5),
   `rep` INT(5),
   `comment` VARCHAR(300),
   `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES  users (id)
        ON DELETE CASCADE ON UPDATE CASCADE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

-- test data
insert into users(name, email) value('user1', 'user1@example.com');
insert into tasks(user_id,menu,set_count,rep,comment) value(1,"benchpress",3,30,"very hard");