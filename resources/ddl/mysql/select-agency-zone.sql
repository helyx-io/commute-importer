select min(s.`stop_lat`), max(s.`stop_lat`), min(s.`stop_lon`), max(s.`stop_lon`) from `%s`.`stops` s