CREATE TABLE %s.trips (
  route_id integer NOT NULL,
  service_id integer NOT NULL,
  trip_id integer NOT NULL,
  trip_headsign varchar(45) DEFAULT NULL,
  direction_id integer DEFAULT NULL,
  block_id varchar(45) DEFAULT NULL,
  shape_id varchar(45) DEFAULT NULL,
  PRIMARY KEY (trip_id)
);
