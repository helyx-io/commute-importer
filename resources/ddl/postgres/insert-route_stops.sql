insert into %s.route_stops (route_id, stop_id, stop_code) select distinct
    r.route_id, s.stop_id, s.stop_code
from
    %s.stops s inner join
    %s.stop_times st on st.stop_id = s.stop_id inner join
    %s.trips t on st.trip_id = t.trip_id inner join
    %s.routes r on t.route_id=r.route_id;
