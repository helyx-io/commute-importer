select st.arrival_time, st.departure_time, st.stop_sequence, s.stop_name from %s.stop_times st inner join %s.stops s on st.stop_id=s.stop_id where st.trip_id='%s' order by st.stop_sequence