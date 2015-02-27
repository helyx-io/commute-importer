CREATE TABLE `gtfs_%s`.`routes` (
  `route_id` int(11) NOT NULL,
  `agency_id` int(11) DEFAULT NULL,
  `route_short_name` varchar(45) DEFAULT NULL,
  `route_long_name` varchar(128) DEFAULT NULL,
  `route_desc` varchar(64) DEFAULT NULL,
  `route_type` int(11) DEFAULT NULL,
  `route_url` varchar(45) DEFAULT NULL,
  `route_color` varchar(45) DEFAULT NULL,
  `route_text_color` varchar(45) DEFAULT NULL,
  -- PRIMARY KEY (`route_id`),
  KEY `route_id_idx` (`route_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
