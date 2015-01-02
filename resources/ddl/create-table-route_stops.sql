CREATE TABLE `gtfs_%s`.`route_stops` (
  `route_id` varchar(45) NOT NULL,
  `stop_id` varchar(45) NOT NULL,
  `stop_code` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`route_id`,`stop_id`),
  KEY `route_id_idx` (`route_id`),
  KEY `stop_id_idx` (`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
