insert into %s.line_stops (line_id, stop_id, stop_code) select distinct
	 l.line_id, s.stop_id, s.stop_code
from
	%s.stops s inner join
    %s.stop_times st on s.stop_id=st.stop_id  inner join
    %s.trips t on t.trip_id=st.trip_id  inner join
    %s.routes r on t.route_id=r.route_id inner join
    %s.lines l on r.route_short_name=l.line_name
order by
    l.line_id, s.stop_id;
