CREATE TABLE %s.calendars (
  service_id integer NOT NULL,
  monday boolean DEFAULT NULL,
  tuesday boolean DEFAULT NULL,
  wednesday boolean DEFAULT NULL,
  thursday boolean DEFAULT NULL,
  friday boolean DEFAULT NULL,
  saturday boolean DEFAULT NULL,
  sunday boolean DEFAULT NULL,
  start_date date DEFAULT NULL,
  end_date date DEFAULT NULL,
  PRIMARY KEY (service_id)
);

