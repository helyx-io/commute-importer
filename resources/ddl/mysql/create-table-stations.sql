CREATE TABLE `%s`.`stations` (
  `station_id` int(11) NOT NULL AUTO_INCREMENT,
  `station_name` varchar(64) NOT NULL,
  `station_lat` float DEFAULT NULL,
  `station_lon` float DEFAULT NULL,
  `station_geo` geometry NOT NULL,
  PRIMARY KEY (`station_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
