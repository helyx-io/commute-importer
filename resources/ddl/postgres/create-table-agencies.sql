CREATE TABLE %s.agencies (
  agency_id integer NOT NULL,
  agency_name char({{agency_name}}) DEFAULT NULL,
  agency_url char({{agency_url}}) DEFAULT NULL,
  agency_timezone char({{agency_timezone}}) DEFAULT NULL,
  agency_lang char({{agency_lang_type}}) DEFAULT NULL,
  agency_min_lat float DEFAULT NULL,
  agency_max_lat float DEFAULT NULL,
  agency_min_lon float DEFAULT NULL,
  agency_max_lon float DEFAULT NULL,
  PRIMARY KEY (agency_id)
);
