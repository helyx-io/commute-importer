CREATE TABLE %s.route_stops (
  route_id integer NOT NULL,
  stop_id integer NOT NULL,
  stop_code char({{length .stops.stop_code}}) NOT NULL,
  PRIMARY KEY (route_id,stop_id)
);
