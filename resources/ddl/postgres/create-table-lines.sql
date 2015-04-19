CREATE TABLE %s.lines (
  line_id SERIAL,
  line_name varchar({{length .routes.route_short_name}}) NOT NULL,
  PRIMARY KEY (line_id)
);
