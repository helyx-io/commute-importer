ALTER TABLE `gtfs_%s`.`stop_times`
    ADD PRIMARY KEY (`trip_id`, `stop_id`),
    ADD INDEX `trip_id_idx` (`trip_id` ASC),
    ADD INDEX `stop_id_idx` (`stop_id` ASC);
