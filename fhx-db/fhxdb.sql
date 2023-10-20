-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db8.hostpark.net
-- Erstellungszeit: 09. Okt 2023 um 09:11
-- Server-Version: 8.0.32
-- PHP-Version: 8.1.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Datenbank: `schmidschluch4`
--
CREATE DATABASE IF NOT EXISTS `schmidschluch4` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `schmidschluch4`;

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `parameters`
--

CREATE TABLE `parameters` (
  `param_id` int NOT NULL,
  `parameter_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `unit_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `parameters`
--

INSERT INTO `parameters` (`param_id`, `parameter_name`, `description`, `created_at`, `updated_at`, `unit_id`) VALUES
(40, 'FP_FSB_PSH', 'OP schliesst ab wenn Druck < Wert', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(41, 'FP_FSB_PSL', 'OP schliesst ab wenn Druck > Wert', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(42, 'FP_OPTION', 'Option auf Ablufsystem oder Entlüftung offen', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(43, 'FP_PAH', 'Druck Alarm HOCH', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(44, 'FP_PAHH', 'Druck Alarm HOCH HOCH [SS]', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(45, 'FP_PAL', 'Druck alarm TIEF', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 174),
(46, 'FP_PALL', 'Druck Alarm TIEF TIEF [SS]', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(47, 'FP_PC_W', 'Druck Sollwert', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(48, 'FP_PV_KENNLINIE', 'Hinterlegte Kennlinie aktivieren', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(49, 'FP_RAMPE', 'Rampe 1 mbar/minute', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(50, 'FP_BESCHREI_TEXT', 'Beschreibung zur OP', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(51, 'FP_PC_W_HO', 'Druck Sollwert bei SS', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174),
(52, 'FP_OPTION_VAK_HO', 'Zustand Vakuum bei einem SS', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 174);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `paramvalues`
--

CREATE TABLE `paramvalues` (
  `value_id` int NOT NULL,
  `stringvalue` varchar(255) DEFAULT NULL,
  `value_set` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `high` int DEFAULT NULL,
  `low` int DEFAULT NULL,
  `cv` int DEFAULT NULL,
  `unit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `param_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `paramvalues`
--

INSERT INTO `paramvalues` (`value_id`, `stringvalue`, `value_set`, `high`, `low`, `cv`, `unit`, `created_at`, `updated_at`, `param_id`) VALUES
(40, '', '', 20000, -1000, 150, '0', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 40),
(41, '', '', 20000, -1000, -150, '0', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 41),
(42, 'ABLUFT1_REGLER', 'LGF_DRUCK_OPT', 0, 0, 0, '0', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 42),
(43, '', '', 20000, -1000, 1150, '0', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 43),
(44, '', '', 20000, -1000, 1200, '0', '2023-08-28 06:34:59', '2023-08-28 06:34:59', 44),
(45, '', '', 20000, -1000, -200, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 45),
(46, '', '', 20000, -1000, -300, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 46),
(47, '', '', 20000, -1000, 50, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 47),
(48, 'AUS', 'L_EIN_AUS', 0, 0, 0, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 48),
(49, '', '', 20000, 0, 0, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 49),
(50, '', '', 0, 0, 0, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 50),
(51, '', '', 20000, -1000, 50, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 51),
(52, 'VAKUUM', 'LGF_DRUCK_VAK_HO', 0, 0, 0, '0', '2023-08-28 06:35:00', '2023-08-28 06:35:00', 52);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `plants`
--

CREATE TABLE `plants` (
  `id` int NOT NULL,
  `plant_name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `plants`
--

INSERT INTO `plants` (`id`, `plant_name`, `created_at`, `updated_at`) VALUES
(1, 'E13', '2023-07-14 09:03:35', '2023-07-14 09:03:35'),
(2, 'D29', '2023-07-14 09:03:35', '2023-07-14 09:03:35'),
(3, 'C07', '2023-07-14 09:05:33', '2023-07-14 09:05:33'),
(4, 'B12', '2023-07-14 09:05:48', '2023-07-14 09:05:48');

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `recipes`
--

CREATE TABLE `recipes` (
  `id` int NOT NULL,
  `recipe_name` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `unit_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `testtabelle`
--

CREATE TABLE `testtabelle` (
  `id` int NOT NULL,
  `name` int NOT NULL,
  `descr` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `testtabelle`
--

INSERT INTO `testtabelle` (`id`, `name`, `descr`) VALUES
(4, 3, 1),
(1, 3, 3),
(2, 3, 4);

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `types`
--

CREATE TABLE `types` (
  `id` int NOT NULL,
  `type_name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `types`
--

INSERT INTO `types` (`id`, `type_name`, `created_at`, `updated_at`) VALUES
(1, 'Operation', '2023-07-12 09:10:07', '2023-07-12 09:10:07'),
(2, 'Unitprocedure', '2023-07-12 09:10:41', '2023-07-12 09:10:41'),
(3, 'Recipe', '2023-07-12 09:10:41', '2023-07-12 09:10:41');

-- --------------------------------------------------------

--
-- Tabellenstruktur für Tabelle `units`
--

CREATE TABLE `units` (
  `id` int NOT NULL,
  `type_id` int NOT NULL,
  `plant_id` int NOT NULL,
  `unit_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `position` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `time` int DEFAULT NULL,
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Daten für Tabelle `units`
--

INSERT INTO `units` (`id`, `type_id`, `plant_id`, `unit_name`, `position`, `time`, `author`, `description`, `created_at`, `updated_at`) VALUES
(174, 1, 2, 'OP_DRUCK', '', 1675757724, 'Hutter Fredy', 'Druck', '2023-08-28 06:34:59', '2023-08-28 06:34:59');

--
-- Tabelle Operation erstellen
--
CREATE VIEW all_operations AS SELECT 
    c.id
    , type_name
    , plant_name
    , description
    , author
FROM units AS c
INNER JOIN types AS a
    ON type_id = a.id
INNER JOIN plants AS b
    ON plant_id = b.id

--
-- Indizes der exportierten Tabellen
--

--
-- Indizes für die Tabelle `parameters`
--
ALTER TABLE `parameters`
  ADD PRIMARY KEY (`param_id`),
  ADD KEY `fk_unit_procedures` (`unit_id`);

--
-- Indizes für die Tabelle `paramvalues`
--
ALTER TABLE `paramvalues`
  ADD PRIMARY KEY (`value_id`),
  ADD KEY `fk_params_value` (`param_id`);

--
-- Indizes für die Tabelle `plants`
--
ALTER TABLE `plants`
  ADD PRIMARY KEY (`id`);

--
-- Indizes für die Tabelle `recipes`
--
ALTER TABLE `recipes`
  ADD PRIMARY KEY (`id`);

--
-- Indizes für die Tabelle `testtabelle`
--
ALTER TABLE `testtabelle`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `zwei` (`name`,`descr`);

--
-- Indizes für die Tabelle `types`
--
ALTER TABLE `types`
  ADD PRIMARY KEY (`id`);

--
-- Indizes für die Tabelle `units`
--
ALTER TABLE `units`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unit_name` (`unit_name`,`plant_id`) USING BTREE;

--
-- AUTO_INCREMENT für exportierte Tabellen
--

--
-- AUTO_INCREMENT für Tabelle `parameters`
--
ALTER TABLE `parameters`
  MODIFY `param_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=53;

--
-- AUTO_INCREMENT für Tabelle `paramvalues`
--
ALTER TABLE `paramvalues`
  MODIFY `value_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=53;

--
-- AUTO_INCREMENT für Tabelle `plants`
--
ALTER TABLE `plants`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT für Tabelle `recipes`
--
ALTER TABLE `recipes`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT für Tabelle `testtabelle`
--
ALTER TABLE `testtabelle`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT für Tabelle `types`
--
ALTER TABLE `types`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT für Tabelle `units`
--
ALTER TABLE `units`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=178;

--
-- Constraints der exportierten Tabellen
--

--
-- Constraints der Tabelle `parameters`
--
ALTER TABLE `parameters`
  ADD CONSTRAINT `fk_unit_procedures` FOREIGN KEY (`unit_id`) REFERENCES `units` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints der Tabelle `paramvalues`
--
ALTER TABLE `paramvalues`
  ADD CONSTRAINT `fk_params_value` FOREIGN KEY (`param_id`) REFERENCES `parameters` (`param_id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
