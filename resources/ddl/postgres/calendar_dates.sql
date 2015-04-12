CREATE TABLE %s.calendar_dates (
  service_id integer NOT NULL,
  date date NOT NULL,
  exception_type integer DEFAULT NULL,
  PRIMARY KEY (service_id, date)
);
