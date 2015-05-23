
--
-- Table structure for table accessTokens
--

DROP TABLE IF EXISTS commute.access_tokens;

CREATE TABLE commute.access_tokens (
  id SERIAL,
  token varchar(255) DEFAULT NULL,
  user_id varchar(255) DEFAULT NULL,
  client_id varchar(255) DEFAULT NULL,
  scope varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

--
-- Table structure for table agencies
--

DROP TABLE IF EXISTS commute.agencies;

CREATE TABLE commute.agencies (
  agency_key varchar(45) NOT NULL,
  agency_id integer NOT NULL,
  agency_name varchar(45) DEFAULT NULL,
  agency_url varchar(45) DEFAULT NULL,
  agency_timezone varchar(45) DEFAULT NULL,
  agency_lang varchar(45) DEFAULT NULL,
  agency_min_lat float DEFAULT NULL,
  agency_max_lat float DEFAULT NULL,
  agency_min_lon float DEFAULT NULL,
  agency_max_lon float DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (agency_key,agency_id)
);


--
-- Table structure for table clients
--

DROP TABLE IF EXISTS commute.clients;

CREATE TABLE commute.clients (
  id SERIAL,
  name varchar(255) DEFAULT NULL,
  client_id varchar(255) DEFAULT NULL,
  client_secret varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

--
-- Table structure for table pem_clients
--

DROP TABLE IF EXISTS commute.pem_clients;

CREATE TABLE commute.pem_clients (
  id SERIAL,
  user_id integer DEFAULT NULL,
  private_key varchar(2048) DEFAULT NULL,
  public_key varchar(2048) DEFAULT NULL,
  certificate varchar(2048) DEFAULT NULL,
  fingerprint varchar(128) DEFAULT NULL,
  key_password varchar(128) DEFAULT NULL,
  cert_password varchar(128) DEFAULT NULL,
  days integer DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

--
-- Table structure for table users
--

-- DROP TABLE IF EXISTS commute.users;
-- CREATE TABLE commute.users (
--  id SERIAL,
--  firstname varchar(255) DEFAULT NULL,
--  lastname varchar(255) DEFAULT NULL,
--  gender varchar(255) DEFAULT NULL,
--  email varchar(255) DEFAULT NULL,
--  google_id varchar(255) DEFAULT NULL,
--  avatar_url varchar(255) DEFAULT NULL,
--  role varchar(255) DEFAULT NULL,
--  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  PRIMARY KEY (id)
-- );

DROP TABLE IF EXISTS commute.users;
CREATE TABLE commute.users (
  id SERIAL,
  firstname varchar(255) DEFAULT NULL,
  lastname varchar(255) DEFAULT NULL,
  gender varchar(255) DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  google_id varchar(255) DEFAULT NULL,
  password varchar(255) DEFAULT NULL,
  role varchar(255) DEFAULT NULL,
  avatar_url varchar(255) DEFAULT NULL,
  token varchar(128) DEFAULT NULL,
  reset_token varchar(128) DEFAULT NULL,
  reset_demand_expiration_date timestamp DEFAULT CURRENT_TIMESTAMP,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (email)
);

--
-- Table structure for table activities
--

DROP TABLE IF EXISTS commute.activities;
CREATE TABLE commute.activities (
  id SERIAL,
  verb varchar(16) DEFAULT NULL,
  published timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  actor_object_type varchar(64) DEFAULT NULL,
  actor_id varchar(128) DEFAULT NULL,
  actor_display_name varchar(128) DEFAULT NULL,
  actor_url varchar(256) DEFAULT NULL,
  actor_image_url varchar(256) DEFAULT NULL,
  actor_image_media_type varchar(32) DEFAULT NULL,
  actor_image_width integer DEFAULT NULL,
  actor_image_height integer DEFAULT NULL,
  object_type varchar(64) DEFAULT NULL,
  object_id varchar(128) DEFAULT NULL,
  object_url varchar(256) DEFAULT NULL,
  object_display_name varchar(128) DEFAULT NULL,
  target_object_type varchar(64) DEFAULT NULL,
  target_id varchar(128) DEFAULT NULL,
  target_display_name varchar(128) DEFAULT NULL,
  target_url varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS commute.sessions;
CREATE TABLE commute.sessions (
  sid varchar NOT NULL COLLATE "default",
  sess json NOT NULL,
  expire timestamp(6) NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE commute.sessions ADD CONSTRAINT session_pkey PRIMARY KEY (sid) NOT DEFERRABLE INITIALLY IMMEDIATE;
