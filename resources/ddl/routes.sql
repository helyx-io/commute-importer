CREATE TABLE `gtfs_%s`.`routes` (
  `route_id` int(11) NOT NULL,
  `agency_id` int(11) NOT NULL,
  `route_short_name` varchar(32) DEFAULT NULL,
  `route_long_name` varchar(128) DEFAULT NULL,
  `route_desc` varchar(64) DEFAULT NULL,
  `route_type` int(11) DEFAULT NULL,
  `route_url` varchar(45) DEFAULT NULL,
  `route_color` char(6) DEFAULT NULL,
  `route_text_color` char(6) DEFAULT NULL,
  PRIMARY KEY (`route_id`),
  KEY `agency_id_idx` (`agency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
