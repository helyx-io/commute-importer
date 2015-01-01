insert into `gtfs_%s`.`stations` (station_name, station_lat, station_lon, station_geo) select distinct
    s.stop_name,
    avg(s.stop_lat),
    avg(s.stop_lon),
    GeomFromText(concat('POINT(', avg(s.stop_lat), ' ', avg(s.stop_lon), ')'))
from
    `gtfs_%s`.`stops` s
group by s.stop_name;
