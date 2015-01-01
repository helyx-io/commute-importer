insert into `lines` (line_name) select distinct r.route_short_name as line_name from routes r order by r.route_short_name;

insert into line_stops (line_id, stop_id, stop_code) select distinct
	 l.line_id, s.stop_id, s.stop_code
from
	stops s inner join
    stop_times st on s.stop_id=st.stop_id  inner join
    trips t on t.trip_id=st.trip_id  inner join
    routes r on t.route_id=r.route_id inner join
    `lines` l on r.route_short_name = l.line_name
order by
    l.line_id, s.stop_id;

insert into stations (station_name, station_lat, station_lon, station_geo)
select distinct s.stop_name, avg(s.stop_lat), avg(s.stop_lon), GeomFromText(concat('POINT(', avg(s.stop_lat), ' ', avg(s.stop_lon), ')'))
from stops s group by s.stop_name;

insert into station_stops (station_id, stop_id) select distinct
    st.station_id, s.stop_id
from
    stops s  inner join
    stations st on s.stop_name = st.station_name;

insert into station_lines (station_id, line_id) select distinct
    s.station_id, ls.line_id
from
    stations s inner join
    station_stops ss on s.station_id=ss.station_id inner join
    line_stops ls on ss.stop_id=ls.stop_id inner join
    `lines` l on l.line_id=ls.line_id;
