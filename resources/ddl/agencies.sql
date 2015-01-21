CREATE TABLE `gtfs_%s`.`agencies` (
  `agency_id` varchar(45) NOT NULL,
  `agency_name` varchar(45) DEFAULT NULL,
  `agency_url` varchar(45) DEFAULT NULL,
  `agency_timezone` varchar(45) DEFAULT NULL,
  `agency_lang` varchar(45) DEFAULT NULL,
  `agency_min_lat` float DEFAULT NULL,
  `agency_max_lat` float DEFAULT NULL,
  `agency_min_lon` float DEFAULT NULL,
  `agency_max_lon` float DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`agency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
