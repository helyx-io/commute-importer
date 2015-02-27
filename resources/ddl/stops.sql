CREATE TABLE `gtfs_%s`.`stops` (
  `stop_id` int(11) NOT NULL,
  `stop_code` varchar(45) DEFAULT NULL,
  `stop_name` varchar(64) DEFAULT NULL,
  `stop_desc` varchar(128) DEFAULT NULL,
  `stop_lat` FLOAT DEFAULT NULL,
  `stop_lon` FLOAT DEFAULT NULL,
  `stop_geo` GEOMETRY DEFAULT NULL,
  `zone_id` varchar(45) DEFAULT NULL,
  `stop_url` varchar(45) DEFAULT NULL,
  `location_type` int(11) DEFAULT NULL,
  `parent_station` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
