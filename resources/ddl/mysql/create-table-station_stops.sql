CREATE TABLE `%s`.`station_stops` (
  `station_id` int(11) NOT NULL,
  `stop_id` int(11) NOT NULL,
  PRIMARY KEY (`station_id`,`stop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
