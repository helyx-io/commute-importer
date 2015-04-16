CREATE TABLE %s.stops (
  stop_id integer NOT NULL,
  stop_code char({{stop_code}}) DEFAULT NULL,
  stop_name char({{stop_name}}) DEFAULT NULL,
  stop_desc varchar({{stop_desc}}) DEFAULT NULL,
  stop_lat FLOAT DEFAULT NULL,
  stop_lon FLOAT DEFAULT NULL,
  stop_geo GEOMETRY DEFAULT NULL,
  zone_id char({{zone_id}}) DEFAULT NULL,
  stop_url char({{stop_url}}) DEFAULT NULL,
  location_type integer DEFAULT NULL,
  parent_station char({{parent_station}}) DEFAULT NULL,
  PRIMARY KEY (stop_id)
);
