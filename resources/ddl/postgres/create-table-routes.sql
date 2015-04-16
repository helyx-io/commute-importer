CREATE TABLE %s.routes (
  route_id integer NOT NULL,
  agency_id integer NOT NULL,
  route_short_name char({{route_short_name}}) DEFAULT NULL,
  route_long_name char({{route_long_name}}) DEFAULT NULL,
  route_desc char({{route_desc}}) DEFAULT NULL,
  route_type integer DEFAULT NULL,
  route_url char({{route_url}}) DEFAULT NULL,
  route_color char(6) DEFAULT NULL,
  route_text_color char(6) DEFAULT NULL,
  PRIMARY KEY (route_id)
);
