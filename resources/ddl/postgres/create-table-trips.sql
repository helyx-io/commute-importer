CREATE TABLE %s.trips (
  route_id integer NOT NULL,
  service_id integer NOT NULL,
  trip_id integer NOT NULL,
  trip_headsign char({{length .trips.trip_headsign}}) DEFAULT NULL,
  direction_id integer DEFAULT NULL,
  block_id char({{length .trips.block_id}}) DEFAULT NULL,
  shape_id char({{length .trips.shape_id}}) DEFAULT NULL,
  PRIMARY KEY (trip_id)
);

