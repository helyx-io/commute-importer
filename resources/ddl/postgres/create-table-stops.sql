CREATE TABLE %s.stops (
  stop_id integer NOT NULL,
  stop_code varchar({{length .stops.stop_code}}) DEFAULT NULL,
  stop_name varchar({{length .stops.stop_name}}) DEFAULT NULL,
  stop_desc varchar({{length .stops.stop_desc}}) DEFAULT NULL,
  stop_lat FLOAT DEFAULT NULL,
  stop_lon FLOAT DEFAULT NULL,
  stop_geo GEOMETRY DEFAULT NULL,
  zone_id varchar({{length .stops.zone_id}}) DEFAULT NULL,
  stop_url varchar({{length .stops.stop_url}}) DEFAULT NULL,
  location_type integer DEFAULT NULL,
  parent_station varchar({{length .stops.parent_station}}) DEFAULT NULL,
  PRIMARY KEY (stop_id)
);
