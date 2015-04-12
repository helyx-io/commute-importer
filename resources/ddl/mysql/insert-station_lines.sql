insert into `%s`.`station_lines` (station_id, line_id) select distinct
    s.station_id, ls.line_id
from
    `%s`.`stations` s inner join
    `%s`.`station_stops` ss on s.station_id=ss.station_id inner join
    `%s`.`line_stops` ls on ss.stop_id=ls.stop_id inner join
    `%s`.`lines` l on l.line_id=ls.line_id;
