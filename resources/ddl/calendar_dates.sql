CREATE TABLE `%s_calendar_dates` (
  `service_id` int(11) NOT NULL,
  `date` date NOT NULL,
  `exception_type` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`service_id`, `date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
