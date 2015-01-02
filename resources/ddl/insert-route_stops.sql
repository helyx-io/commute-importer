insert into `gtfs_%s`.`route_stops` (route_id, stop_id, stop_code) select distinct
    r.route_id, s.stop_id, s.stop_code
from
    `gtfs_%s`.`stops` s inner join
    `gtfs_%s`.`stop_times` st on st.stop_id = s.stop_id inner join
    `gtfs_%s`.`trips` t on st.trip_id = t.trip_id inner join
    `gtfs_%s`.`routes` r on t.route_id=r.route_id;
