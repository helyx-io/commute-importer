CREATE TABLE %s.line_stops (
  line_id integer NOT NULL,
  stop_id integer NOT NULL,
  stop_code varchar(45) NOT NULL,
  PRIMARY KEY (line_id,stop_id)
);
