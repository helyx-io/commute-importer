CREATE TABLE `%s`.`line_stops` (
  `line_id` int(11) NOT NULL,
  `stop_id` int(11) NOT NULL,
  `stop_code` varchar(45) NOT NULL,
  PRIMARY KEY (`line_id`,`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
