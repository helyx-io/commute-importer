CREATE TABLE `gtfs_%s`.`station_lines` (
  `station_id` integer(11) NOT NULL,
  `line_id` integer(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`station_id`, `line_id`),
  KEY `station_id_idx` (`station_id`),
  KEY `line_id_idx` (`line_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
