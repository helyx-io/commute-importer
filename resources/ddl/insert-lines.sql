insert into `gtfs_%s`.`lines` (line_name) select distinct
    r.route_short_name as line_name
from
    `gtfs_%s`.`routes` r
order by
    r.route_short_name;
