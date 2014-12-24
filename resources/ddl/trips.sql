CREATE TABLE `%s_trips` (
  `route_id` varchar(45) DEFAULT NULL,
  `service_id` int(11) DEFAULT NULL,
  `trip_id` varchar(45) NOT NULL,
  `trip_headsign` varchar(45) DEFAULT NULL,
  `direction_id` int(11) DEFAULT NULL,
  `block_id` varchar(45) DEFAULT NULL,
  `shape_id` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`trip_id`),
  KEY `route_id_idx` (`route_id`),
  KEY `service_id_idx` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
