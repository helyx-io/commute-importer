CREATE TABLE `gtfs_%s`.`stations` (
  `station_id` integer(11) NOT NULL AUTO_INCREMENT,
  `station_name` varchar(64) NOT NULL,
  `station_lat` float DEFAULT NULL,
  `station_lon` float DEFAULT NULL,
  `station_geo` geometry NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`station_id`),
  KEY `station_name_idx` (`station_name`),
  SPATIAL KEY `station_geo_idx` (`station_geo`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
