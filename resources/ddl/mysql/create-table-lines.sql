CREATE TABLE `%s`.`lines` (
  `line_id` int(11) NOT NULL AUTO_INCREMENT,
  `line_name` varchar(45) NOT NULL,
  PRIMARY KEY (`line_id`),
  KEY `line_name_idx` (`line_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
