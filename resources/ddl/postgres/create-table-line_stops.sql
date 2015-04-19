CREATE TABLE %s.line_stops (
  line_id integer NOT NULL,
  stop_id integer NOT NULL,
  stop_code varchar({{length .stops.stop_code}}) NOT NULL,
  PRIMARY KEY (line_id,stop_id)
);
