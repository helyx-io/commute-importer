CREATE TABLE %s.route_stops (
  route_id integer NOT NULL,
  stop_id integer NOT NULL,
  stop_code varchar(45) NOT NULL,
  PRIMARY KEY (route_id,stop_id)
);
