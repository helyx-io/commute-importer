CREATE TABLE `gtfs_%s`.`station_stops` (
  `station_id` integer(11) NOT NULL,
  `stop_id` varchar(64) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`station_id`,`stop_id`),
  KEY `station_id_idx` (`station_id`),
  KEY `stop_id_idx` (`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
