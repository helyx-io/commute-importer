CREATE TABLE `%s`.`transfers` (
  `from_stop_id` varchar(45) NOT NULL,
  `to_stop_id` varchar(45) NOT NULL,
  `transfer_type` int(11) NOT NULL DEFAULT '0',
  `min_transfer_time` int(11) DEFAULT NULL --,
  -- PRIMARY KEY (`from_stop_id`,`to_stop_id`,`transfer_type`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
