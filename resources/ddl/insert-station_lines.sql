insert into `gtfs_%s`.`station_lines` (station_id, line_id) select distinct
    s.station_id, ls.line_id
from
    `gtfs_%s`.`stations` s inner join
    `gtfs_%s`.`station_stops` ss on s.station_id=ss.station_id inner join
    `gtfs_%s`.`line_stops` ls on ss.stop_id=ls.stop_id inner join
    `gtfs_%s`.`lines` l on l.line_id=ls.line_id;
