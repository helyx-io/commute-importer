CREATE TABLE %s.routes (
  route_id integer NOT NULL,
  agency_id integer NOT NULL,
  route_short_name varchar({{length .routes.route_short_name}}) DEFAULT NULL,
  route_long_name varchar({{length .routes.route_long_name}}) DEFAULT NULL,
  route_desc varchar({{length .routes.route_desc}}) DEFAULT NULL,
  route_type integer DEFAULT NULL,
  route_url varchar({{length .routes.route_url}}) DEFAULT NULL,
  route_color integer DEFAULT NULL,
  route_text_color integer DEFAULT NULL,
  PRIMARY KEY (route_id)
);
