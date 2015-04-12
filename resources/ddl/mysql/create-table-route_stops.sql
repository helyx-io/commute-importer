CREATE TABLE `%s`.`route_stops` (
  `route_id` int(11) NOT NULL,
  `stop_id` int(11) NOT NULL,
  `stop_code` varchar(45) NOT NULL,
  PRIMARY KEY (`route_id`,`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
