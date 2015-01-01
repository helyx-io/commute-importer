insert into `gtfs_%s`.`line_stops` (line_id, stop_id, stop_code) select distinct
	 l.line_id, s.stop_id, s.stop_code
from
	`gtfs_%s`.`stops` s inner join
    `gtfs_%s`.`stop_times` st on s.stop_id=st.stop_id  inner join
    `gtfs_%s`.`trips` t on t.trip_id=st.trip_id  inner join
    `gtfs_%s`.`routes` r on t.route_id=r.route_id inner join
    `gtfs_%s`.`lines` l on r.route_short_name=l.line_name
order by
    l.line_id, s.stop_id;
