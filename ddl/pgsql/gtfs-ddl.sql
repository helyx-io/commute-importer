--
-- Table structure for table accessTokens
--

DROP TABLE IF EXISTS gtfs.access_tokens;

CREATE TABLE gtfs.access_tokens (
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

DROP TABLE IF EXISTS gtfs.agencies;

CREATE TABLE gtfs.agencies (
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

DROP TABLE IF EXISTS gtfs.clients;

CREATE TABLE gtfs.clients (
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

DROP TABLE IF EXISTS gtfs.pem_clients;

CREATE TABLE gtfs.pem_clients (
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

DROP TABLE IF EXISTS gtfs.users;
CREATE TABLE gtfs.users (
  id SERIAL,
  firstName varchar(255) DEFAULT NULL,
  lastName varchar(255) DEFAULT NULL,
  gender varchar(255) DEFAULT NULL,
  email varchar(255) DEFAULT NULL,
  googleId varchar(255) DEFAULT NULL,
  avatarUrl varchar(255) DEFAULT NULL,
  role varchar(255) DEFAULT NULL,
  createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
