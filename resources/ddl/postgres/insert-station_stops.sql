insert into %s.station_stops (station_id, stop_id) select distinct
    st.station_id, s.stop_id
from
    %s.stops s  inner join
    %s.stations st on s.stop_name = st.station_name;
