CREATE TABLE `gtfs_%s`.`lines` (
  `line_id` integer(11) NOT NULL AUTO_INCREMENT,
  `line_name` varchar(45) NOT NULL,
  -- PRIMARY KEY (`line_id`),
  KEY `line_id_idx` (`line_id`),
  KEY `line_name_idx` (`line_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
