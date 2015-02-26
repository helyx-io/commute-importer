CREATE TABLE `gtfs_%s`.`line_stops` (
  `line_id` integer(11) NOT NULL,
  `stop_id` varchar(64) NOT NULL,
  `stop_code` varchar(45) NOT NULL,
  -- PRIMARY KEY (`line_id`,`stop_id`)
  KEY `line_id_idx` (`line_id`),
  KEY `stop_id_idx` (`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
