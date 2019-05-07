SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;


CREATE TABLE IF NOT EXISTS `buses`
(
    `id`          int(10) UNSIGNED                       NOT NULL AUTO_INCREMENT,
    `license`     varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
    `total_seats` int(5)                                 NOT NULL,
    `empty_seats` int(5)                                 NOT NULL,
    `departure`   varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `destination` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `begin_at`    time                                   NOT NULL,
    `end_at`      time                                   NOT NULL,
    `price`       double                                 NOT NULL DEFAULT 0,
    `info`        text COLLATE utf8mb4_unicode_ci                 DEFAULT NULL,
    `weekly`      tinyint(3)                             NOT NULL DEFAULT 0,
    `status`      tinyint(1)                             NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `codes`
(
    `id`      int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `code`    varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `status`  tinyint(1)       NOT NULL              DEFAULT -1,
    `user_id` int(10) UNSIGNED NOT NULL,
    `useAt`   datetime         NOT NULL              DEFAULT current_timestamp(),
    PRIMARY KEY (`id`),
    UNIQUE KEY `key` (`code`),
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `comments`
(
    `id`              int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content`         text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `comment_at`      datetime         NOT NULL       DEFAULT current_timestamp(),
    `status`          tinyint(1)       NOT NULL       DEFAULT 0,
    `stars`           tinyint(1)                      DEFAULT 5,
    `is_replied`      tinyint(1)       NOT NULL       DEFAULT 0,
    `content_replied` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `user_id`         int(10) UNSIGNED NOT NULL,
    `bus_id`          int(10) UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `bus_id` (`bus_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `favorites`
(
    `id`      int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(10) UNSIGNED NOT NULL,
    `bus_id`  int(10) UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `bus_id` (`bus_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `orders`
(
    `id`       int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_at` datetime         NOT NULL DEFAULT current_timestamp(),
    `status`   tinyint(1)       NOT NULL DEFAULT 0,
    `user_id`  int(10) UNSIGNED NOT NULL,
    `bus_id`   int(10) UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `orders_ibfk_2` (`bus_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `users`
(
    `id`       int(10) UNSIGNED                       NOT NULL AUTO_INCREMENT,
    `account`  varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL,
    `salt`     varchar(10) COLLATE utf8mb4_unicode_ci          DEFAULT NULL,
    `balance`  double                                 NOT NULL DEFAULT 0,
    `is_admin` tinyint(1)                             NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `account` (`account`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;


ALTER TABLE `codes`
    ADD CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `comments`
    ADD CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    ADD CONSTRAINT `comments_ibfk_2` FOREIGN KEY (`bus_id`) REFERENCES `buses` (`id`);

ALTER TABLE `favorites`
    ADD CONSTRAINT `favorites_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    ADD CONSTRAINT `favorites_ibfk_2` FOREIGN KEY (`bus_id`) REFERENCES `buses` (`id`);

ALTER TABLE `orders`
    ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    ADD CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`bus_id`) REFERENCES `buses` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
