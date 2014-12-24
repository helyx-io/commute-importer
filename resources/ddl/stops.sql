CREATE TABLE `gtfs_%s`.`stops` (
  `stop_id` varchar(45) NOT NULL,
  `stop_code` varchar(45) DEFAULT NULL,
  `stop_name` varchar(64) DEFAULT NULL,
  `stop_desc` varchar(128) DEFAULT NULL,
  `stop_lat` int(11) DEFAULT NULL,
  `stop_lon` int(11) DEFAULT NULL,
  `zone_id` varchar(45) DEFAULT NULL,
  `stop_url` varchar(45) DEFAULT NULL,
  `location_type` int(11) DEFAULT NULL,
  `parent_station` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
