/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `records`;
CREATE TABLE `records` (
  `id` varchar(255) NOT NULL,
  `data` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `records` (`id`, `data`, `created_at`) VALUES
('123sdah2f3h23', 'Ungart\'s qr code', '2020-11-25 13:34:51');
INSERT INTO `records` (`id`, `data`, `created_at`) VALUES
('1klcv7vLqgdl0WZlqhAId8shKMC', 'Jako\'s QR code', '2020-11-25 12:34:05');
INSERT INTO `records` (`id`, `data`, `created_at`) VALUES
('1kldAZqJNL5WKjqzLHabAGHGdE5', 'Gary\'s qr code', '2020-11-25 12:36:08');
INSERT INTO `records` (`id`, `data`, `created_at`) VALUES
('1kljrgpwMShoOsEmIVycKE8p4zN', 'Ungart\'s qr code', '2020-11-25 13:31:12'),
('1klLhwJK4FeOqgZDglyz0socRnA', 'Idol Mitoy lng malakas', '2020-11-25 10:12:33'),
('1klLIPLz8jLHjI8BgetQo2jvYcq', 'chuchu', '2020-11-25 10:09:11'),
('1klLn5yjMIlmeuhbrAePjkzQOMz', 'Mitoy lng malakas', '2020-11-25 10:13:14'),
('1klOmgHODB7msn4e6UoPJsOu5Dc', 'test api', '2020-11-25 10:37:51'),
('1klWniaeY9gwfLAgN4hug31Ii3a', 'Test generate qr id', '2020-11-25 11:43:46'),
('1klWZNowQhEQ4IjTZoFikk9OYmQ', 'Test qr generator', '2020-11-25 11:41:52'),
('1klZ1Q7xsDWwsrzY2WtEfNpRwKm', 'I can do all things..', '2020-11-25 12:02:02'),
('1klZ2ptgrknVWnQf92XE84Hq5kO', 'Philippians 4:13', '2020-11-25 12:02:13'),
('213', 'chuchu', '2020-11-25 08:52:28'),
('213123', 'chuchu', '2020-11-25 10:10:40'),
('23y18edidh7q2du', 'Ungart\'s qr code', '2020-11-25 13:35:03'),
('ajshduiwq', 'Mitoy lng malakas', '2020-11-25 10:13:45');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;