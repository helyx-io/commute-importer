CREATE TABLE `%s`.`station_lines` (
  `station_id` int(11) NOT NULL,
  `line_id` int(11) NOT NULL,
  PRIMARY KEY (`station_id`, `line_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
