CREATE TABLE %s.stop_times (
  trip_id integer NOT NULL,
  arrival_time interval DEFAULT NULL,
  departure_time interval DEFAULT NULL,
  stop_id integer NOT NULL,
  stop_sequence integer DEFAULT NULL,
  stop_head_sign varchar({{length .stop_times.stop_head_sign}}) DEFAULT NULL,
  pickup_type integer DEFAULT NULL,
  drop_off_type integer DEFAULT NULL
);
