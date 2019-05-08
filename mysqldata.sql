-- MySQL dump 10.16  Distrib 10.1.37-MariaDB, for Win32 (AMD64)
--
-- Host: localhost    Database: mysqldata
-- ------------------------------------------------------
-- Server version	10.1.37-MariaDB

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

--
-- Table structure for table `tabelcount`
--

DROP TABLE IF EXISTS `tabelcount`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tabelcount` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `countproduct` int(11) NOT NULL,
  `countproducttemp` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tabelcount`
--

LOCK TABLES `tabelcount` WRITE;
/*!40000 ALTER TABLE `tabelcount` DISABLE KEYS */;
INSERT INTO `tabelcount` VALUES (1,15,0);
/*!40000 ALTER TABLE `tabelcount` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tabeldefaultpage`
--

DROP TABLE IF EXISTS `tabeldefaultpage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tabeldefaultpage` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `defaultpage` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tabeldefaultpage`
--

LOCK TABLES `tabeldefaultpage` WRITE;
/*!40000 ALTER TABLE `tabeldefaultpage` DISABLE KEYS */;
INSERT INTO `tabeldefaultpage` VALUES (1,12);
/*!40000 ALTER TABLE `tabeldefaultpage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tabelproduct`
--

DROP TABLE IF EXISTS `tabelproduct`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tabelproduct` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `imagepath` varchar(50) NOT NULL,
  `title` varchar(50) NOT NULL,
  `description` varchar(66) NOT NULL,
  `price` int(11) NOT NULL,
  `idtemp` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tabelproduct_desc_index` (`description`),
  KEY `tabelproduct_id_index` (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tabelproduct`
--

LOCK TABLES `tabelproduct` WRITE;
/*!40000 ALTER TABLE `tabelproduct` DISABLE KEYS */;
INSERT INTO `tabelproduct` VALUES (1,'/images1.jpg','LazMall','CANON Pixma G3010',4500000,0),(2,'/images2.jpg','LazMall','Laptop Asus I5 4Gb',14500000,0),(3,'/images3.jpg','LazMall','Handphone Xiomi',5500000,0),(4,'/images4.jpg','LazMall','Sony Original NW-A56HN/NWA56HN/NW A56HN',3500000,0),(5,'/images5.jpg','LazMall','Logam Bluetooth MP3 Pemutar Musik',6000000,0),(6,'/images6.jpg','LazMall','Asus M32AD-ID008D PC Core i5 1TB',4750000,0),(7,'/images7.jpg','LazMall','Digital Alliance PC Gaming 8100 S',9500000,0),(8,'/images8.jpg','LazMall','Home Theater Samsung',11500000,0),(9,'/images9.jpg','LazMall','Play Station 4',12500000,0),(10,'/images10.jpg','LazMall','Televisi Samsung 50',13500000,0),(11,'/images11.jpg','LazMall','AQUA KULKAS 2PINTU AQR-D275',7500000,0),(12,'/images12.jpg','LazMall','Sofa Bed Valencia - Navy Kain',10500000,0),(13,'/images13.jpg','LazMall','Desktop Lenovo All In One i7-3770S CPU @ 3.10GHz',10500000,0),(14,'/images14.jpg','LazMall','AC Split ½ PK SAMSUNG AS-05TULN',3500000,0),(15,'/images15.jpg','LazMall','Microwave Elecrolux EMS2348',10500000,0);
/*!40000 ALTER TABLE `tabelproduct` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tabelresume`
--

DROP TABLE IF EXISTS `tabelresume`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tabelresume` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL,
  `title` varchar(50) NOT NULL,
  `photo` varchar(40) NOT NULL,
  `description` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tabelresume`
--

LOCK TABLES `tabelresume` WRITE;
/*!40000 ALTER TABLE `tabelresume` DISABLE KEYS */;
INSERT INTO `tabelresume` VALUES (1,'Suharto','President of the Indonesia Republic (2st)','/foto1.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(2,'Vladimir Putin','President of the Russian Republic','/foto2.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(3,'Barack Obama','President of the United States Republic (44th)','/foto3.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(4,'Kim Jong Un','President of the North Korea Republic','/foto4.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(5,'Xi Jinping','President of the People\'s China Republic','/foto5.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(6,'Recep Tayyip Erdogan','President of the Turkey Republic','/foto6.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(7,'Erna Solberg Wesenberg','Prime Minister of Norway','/foto7.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(8,'Angela Merkel','Chancellor of Germany','/foto8.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(9,'Justin Trudeau','Prime Minister of Canada','/foto9.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(10,'Theresa May','Prime Minister of British (past)','/foto10.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(11,'Hassanal Bolkiah','Prime Minister of Brunei Darussalam','/foto11.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.'),(12,'Emmanuel Macron','President of the French Republic','/foto12.jpg','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.');
/*!40000 ALTER TABLE `tabelresume` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tabeluser`
--

DROP TABLE IF EXISTS `tabeluser`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tabeluser` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(35) NOT NULL,
  `password` varchar(100) NOT NULL,
  `level` varchar(10) NOT NULL,
  `email` varchar(60) NOT NULL,
  `notelpon` varchar(15) NOT NULL,
  `token` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10151 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tabeluser`
--

LOCK TABLES `tabeluser` WRITE;
/*!40000 ALTER TABLE `tabeluser` DISABLE KEYS */;
INSERT INTO `tabeluser` VALUES (1051,'user01','$2b$08$2Ml6U3cQvldvWmrhKxuiReXIKDce/cjQzwCyfWnYBLKEPM4VzzIhK','admin','rakuti1965@gmail.com','087880800827','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJqZWN0IjoiIiwiaWF0IjoxNTU2MjI2MzY3fQ.PRNkiF4iMiL0U6I4P4uHaxU44p6SrD_j0PaleYkP0ls');
/*!40000 ALTER TABLE `tabeluser` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'mysqldata'
--
/*!50003 DROP PROCEDURE IF EXISTS `copy_tabelproduct` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `copy_tabelproduct`(wf_searchname varchar(50))
BEGIN
      SET @_wf_search_name = wf_searchname;
      CREATE TABLE IF NOT EXISTS `producttemp` (idtemp INT(11) auto_increment primary key) ENGINE=MYISAM SELECT id,imagepath,title,description,price FROM tabelproduct
      where description LIKE CONCAT('%', @_wf_search_name , '%') ;
      ALTER TABLE producttemp MODIFY idtemp Int(11) AFTER price;
   END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `count_tabelproduct` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `count_tabelproduct`(wf_searchname varchar(50))
BEGIN
      SET @_wf_search_name = wf_searchname;
      SELECT COUNT(*) FROM tabelproduct WHERE `description` LIKE CONCAT('%', @_wf_search_name, '%') ORDER BY ID;
   END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `get_tabelproduct` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `get_tabelproduct`(IN `wf_searchname` VARCHAR(50), IN `wf_limit` INT, IN `wf_offset` INT)
BEGIN
      SET @_wf_search_name = wf_searchname;
      SET @_wf_limit = wf_limit;
      SET @_wf_offset = wf_offset;
      PREPARE stmt FROM "select * from tabelproduct WHERE `description` LIKE CONCAT('%', @_wf_search_name , '%') ORDER BY id LIMIT ? OFFSET ?;";
      EXECUTE stmt USING @_wf_limit, @_wf_offset;
      DEALLOCATE PREPARE stmt;
  END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `update_countproduct` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `update_countproduct`(wf_count int)
BEGIN
      SET @_wf_count = wf_count;
      UPDATE tabelcount SET countproduct= countproduct + @_wf_count;
   END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `update_countproducttemp` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_unicode_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `update_countproducttemp`(wf_count int)
BEGIN
      SET @_wf_count = wf_count;
      UPDATE tabelcount SET countproducttemp= @_wf_count;
   END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-05  3:20:12
