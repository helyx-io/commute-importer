insert into `%s`.`lines` (line_name) select distinct
    r.route_short_name as line_name
from
    `%s`.`routes` r
order by
    r.route_short_name;
