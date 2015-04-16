CREATE TABLE %s.routes (
  route_id integer NOT NULL,
  agency_id integer NOT NULL,
  route_short_name char({{length .routes.route_short_name}}) DEFAULT NULL,
  route_long_name char({{length .routes.route_long_name}}) DEFAULT NULL,
  route_desc char({{length .routes.route_desc}}) DEFAULT NULL,
  route_type integer DEFAULT NULL,
  route_url char({{length .routes.route_url}}) DEFAULT NULL,
  route_color char(6) DEFAULT NULL,
  route_text_color char(6) DEFAULT NULL,
  PRIMARY KEY (route_id)
);
