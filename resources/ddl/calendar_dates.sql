CREATE TABLE `gtfs_%s`.`calendar_dates` (
  `service_id` int(11) NOT NULL,
  `date` date NOT NULL,
  `exception_type` int(11) DEFAULT NULL,
  PRIMARY KEY (`service_id`, `date`),
  KEY `date_idx` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
