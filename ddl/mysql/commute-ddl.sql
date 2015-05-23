-- SQL dump 10.13  Distrib 5.6.15, for osx10.7 (x86_64)
--
-- Host: localhost    Database: commute
-- ------------------------------------------------------
-- Server version	5.6.15

DROP SCHEMA IF EXISTS `commute`;
CREATE SCHEMA `commute` DEFAULT CHARACTER SET utf8mb4;

-- CREATE USER 'commute'@'localhost' IDENTIFIED BY 'commute';
-- GRANT ALL PRIVILEGES ON *.* TO 'commute'@'localhost';
-- CREATE USER 'commute'@'172.17.1.%' IDENTIFIED BY 'commute';
-- GRANT ALL PRIVILEGES ON *.* TO 'commute'@'172.17.1.%';
-- CREATE USER 'commute'@'%' IDENTIFIED BY 'commute';
-- GRANT ALL PRIVILEGES ON *.* TO 'commute'@'%';
-- flush privileges;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

use commute;

--
-- Table structure for table `access_tokens`
--

DROP TABLE IF EXISTS `access_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `access_tokens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `token` varchar(255) DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `client_id` varchar(255) DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `agencies`
--

DROP TABLE IF EXISTS `agencies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `agencies` (
  `agency_key` varchar(45) NOT NULL,
  `agency_id` int(11) NOT NULL,
  `agency_name` varchar(45) DEFAULT NULL,
  `agency_url` varchar(45) DEFAULT NULL,
  `agency_timezone` varchar(45) DEFAULT NULL,
  `agency_lang` varchar(45) DEFAULT NULL,
  `agency_min_lat` float DEFAULT NULL,
  `agency_max_lat` float DEFAULT NULL,
  `agency_min_lon` float DEFAULT NULL,
  `agency_max_lon` float DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`agency_key`,`agency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `client_id` varchar(255) DEFAULT NULL,
  `client_secret` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `pem_clients`
--

DROP TABLE IF EXISTS `pem_clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pem_clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userId` int(11) DEFAULT NULL,
  `private_key` varchar(2048) DEFAULT NULL,
  `public_key` varchar(2048) DEFAULT NULL,
  `certificate` varchar(2048) DEFAULT NULL,
  `fingerprint` varchar(128) DEFAULT NULL,
  `key_password` varchar(128) DEFAULT NULL,
  `cert_password` varchar(128) DEFAULT NULL,
  `days` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

-- DROP TABLE IF EXISTS `users`;
-- /*!40101 SET @saved_cs_client     = @@character_set_client */;
-- /*!40101 SET character_set_client = utf8 */;
-- CREATE TABLE `users` (
--  `id` int(11) NOT NULL AUTO_INCREMENT,
--  `firstname` varchar(255) DEFAULT NULL,
--  `lastname` varchar(255) DEFAULT NULL,
--  `gender` varchar(255) DEFAULT NULL,
--  `email` varchar(255) DEFAULT NULL,
--  `google_id` varchar(255) DEFAULT NULL,
--  `avatar_url` varchar(255) DEFAULT NULL,
--  `role` varchar(255) DEFAULT NULL,
--  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
--  PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- /*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `google_id` varchar(255) DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `token` varchar(128) DEFAULT NULL,
  `reset_token` varchar(128) DEFAULT NULL,
  `reset_demand_expiration_date` timestamp DEFAULT CURRENT_TIMESTAMP,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `activities`
--

DROP TABLE IF EXISTS `activities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `activities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `verb` varchar(16) DEFAULT NULL,
  `published` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `actor_object_type` varchar(64) DEFAULT NULL,
  `actor_id` varchar(128) DEFAULT NULL,
  `actor_display_name` varchar(128) DEFAULT NULL,
  `actor_url` varchar(256) DEFAULT NULL,
  `actor_image_url` varchar(256) DEFAULT NULL,
  `actor_image_media_type` varchar(32) DEFAULT NULL,
  `actor_image_width` int(11) DEFAULT NULL,
  `actor_image_height` int(11) DEFAULT NULL,
  `object_type` varchar(64) DEFAULT NULL,
  `object_id` varchar(128) DEFAULT NULL,
  `object_url` varchar(256) DEFAULT NULL,
  `object_display_name` varchar(128) DEFAULT NULL,
  `target_object_type` varchar(64) DEFAULT NULL,
  `target_id` varchar(128) DEFAULT NULL,
  `target_display_name` varchar(128) DEFAULT NULL,
  `target_url` varchar(256) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-12-31 16:43:35
