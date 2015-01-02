insert into `gtfs_%s`.`stop_times_full`
select 
    s.stop_id,
    s.stop_code,
    s.stop_name,
    s.stop_desc,
    s.stop_lat,
    s.stop_lon,
    s.stop_geo,
    s.zone_id,
    s.stop_url,
    s.location_type,
    s.parent_station,
	st.arrival_time,
    st.departure_time,
    st.stop_sequence,
    st.stop_head_sign,
    st.pickup_type,
    st.drop_off_type,
    t.service_id,
    t.trip_id,
	t.trip_headsign,
    t.direction_id,
    t.block_id,
    t.shape_id,
    r.agency_id,
    r.route_id,
	r.route_short_name,
    r.route_long_name,
    r.route_desc,
    r.route_type,
    r.route_url,
    r.route_color,
    r.route_text_color,
    s.created_at,
    s.updated_at
from 
	`gtfs_%s`.`stops` s inner join
    `gtfs_%s`.`stop_times` st on s.stop_id=st.stop_id inner join
    `gtfs_%s`.`trips` t on st.trip_id=t.trip_id inner join
    `gtfs_%s`.`routes` r on r.route_id=t.route_id
where
	r.route_short_name='%v'