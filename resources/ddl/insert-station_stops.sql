insert into `gtfs_%s`.`station_stops` (station_id, stop_id) select distinct
    st.station_id, s.stop_id
from
    `gtfs_%s`.`stops` s  inner join
    `gtfs_%s`.`stations` st on s.stop_name = st.station_name;
