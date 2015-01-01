CREATE TABLE `gtfs_%s`.`lines` (
  `line_id` integer(11) NOT NULL AUTO_INCREMENT,
  `line_name` varchar(45) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`line_id`),
  KEY `line_name_idx` (`line_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
