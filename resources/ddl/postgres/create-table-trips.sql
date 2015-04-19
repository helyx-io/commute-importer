CREATE TABLE %s.trips (
  route_id integer NOT NULL,
  service_id integer NOT NULL,
  trip_id integer NOT NULL,
  trip_headsign varchar({{length .trips.trip_headsign}}) DEFAULT NULL,
  direction_id integer DEFAULT NULL,
  block_id varchar({{length .trips.block_id}}) DEFAULT NULL,
  shape_id varchar({{length .trips.shape_id}}) DEFAULT NULL,
  PRIMARY KEY (trip_id)
);

