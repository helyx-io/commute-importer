CREATE TABLE `gtfs_%s`.`agencies` (
  `agency_id` int(11) NOT NULL,
  `agency_name` varchar(45) DEFAULT NULL,
  `agency_url` varchar(45) DEFAULT NULL,
  `agency_timezone` varchar(45) DEFAULT NULL,
  `agency_lang` varchar(45) DEFAULT NULL,
  `agency_min_lat` float DEFAULT NULL,
  `agency_max_lat` float DEFAULT NULL,
  `agency_min_lon` float DEFAULT NULL,
  `agency_max_lon` float DEFAULT NULL,
  PRIMARY KEY (`agency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
