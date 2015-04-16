CREATE TABLE %s.stations (
  station_id SERIAL,
  station_name char({{length .stops.stop_name}}) NOT NULL,
  station_lat float DEFAULT NULL,
  station_lon float DEFAULT NULL,
  station_geo geometry NOT NULL,
  PRIMARY KEY (station_id)
);

