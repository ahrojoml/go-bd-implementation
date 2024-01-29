-- DDL: Data Definition Language
DROP DATABASE IF EXISTS `supermarket_db`;

CREATE DATABASE `supermarket_db`;

USE `supermarket_db`;

DROP TABLE IF EXISTS `warehouses`;

CREATE TABLE `warehouses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `address` varchar(150) NOT NULL,
  `telephone` varchar(150) NOT NULL,
  `capacity` int NOT NULL,
  PRIMARY KEY (`id`)
);

-- Volcado de datos para la tabla `warehouses`
INSERT INTO `warehouses` (`name`, `address`, `telephone`, `capacity`) VALUES
('Main Warehouse', '221 Baker Street', "4555666", 100);

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `quantity` int NOT NULL,
  `code_value` varchar(255) NOT NULL,
  `is_published` boolean NOT NULL,
  `expiration` date NOT NULL,
  `price` decimal(10, 2) NOT NULL,
  `warehouse_id` int NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `products_code_value_unique` UNIQUE (`code_value`),
  CONSTRAINT `fk_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO `products` (`name`, `quantity`, `code_value`, `is_published`, `expiration`, `price`, `warehouse_id`) VALUES
("Tomato", 150, "F1001", true, "2024-03-23", 10, 1),
("Potato", 10, "F1002", true, "2024-10-12", 20, 1),
("Onion", 200, "F1003", false, "2024-05-06", 25, 1),
("Carrot", 300, "F1004", true, "2025-01-01", 23, 1);
