# ************************************************************
# Sequel Pro SQL dump
# Version 4468
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.19)
# Database: godmin
# Generation Time: 2018-09-09 15:03:30 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table goadmin_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_menu`;

CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;

INSERT INTO `goadmin_menu` (`id`, `parent_id`, `order`, `title`, `icon`, `uri`, `created_at`, `updated_at`, `type`)
VALUES
	(2,0,2,'Admin','fa-tasks','',NULL,NULL,1),
	(3,2,2,'Users','fa-users','/info/manager',NULL,NULL,1),
	(4,2,3,'Roles','fa-user','/info/roles',NULL,NULL,1),
	(5,2,4,'Permission','fa-ban','/info/permission',NULL,NULL,1),
	(6,2,5,'Menu','fa-bars','/menu',NULL,NULL,1),
	(7,2,6,'Operation log','fa-history','/info/op',NULL,NULL,1),
	(12,0,7,'Users','fa-user','/info/user',NULL,NULL,0),
	(13,0,1,'Dashboard','fa-bar-chart','/','2018-08-03 15:24:42',NULL,1),
	(16,0,8,'Posts','fa-file-powerpoint-o','/info/posts','2018-09-08 08:43:57','2018-09-08 08:43:57',0),
	(17,0,9,'Authors','fa-users','/info/authors','2018-09-08 08:44:38','2018-09-08 08:44:38',0),
	(18,0,10,'Example Plugin','fa-plug','/example','2018-09-08 09:06:41','2018-09-08 09:06:41',0);

/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_operation_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_operation_log`;

CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_operation_log` WRITE;
/*!40000 ALTER TABLE `goadmin_operation_log` DISABLE KEYS */;

INSERT INTO `goadmin_operation_log` (`id`, `user_id`, `path`, `method`, `ip`, `input`, `created_at`, `updated_at`)
VALUES
	(16,1,'admin/auth/menu','POST','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u4e5d\\u56fe\",\"icon\":\"fa-bars\",\"uri\":\"\\/ninepic\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\"}','2018-05-13 10:12:31','2018-05-13 10:12:31'),
	(17,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:12:31','2018-05-13 10:12:31'),
	(18,1,'admin/auth/menu','POST','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u7528\\u6237\\u7ba1\\u7406\",\"icon\":\"fa-bars\",\"uri\":\"\\/user\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\"}','2018-05-13 10:12:44','2018-05-13 10:12:44'),
	(19,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:12:45','2018-05-13 10:12:45'),
	(20,1,'admin/auth/menu','POST','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u914d\\u7f6e\\u7ba1\\u7406\",\"icon\":\"fa-bars\",\"uri\":\"\\/config\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\"}','2018-05-13 10:13:07','2018-05-13 10:13:07'),
	(21,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:13:07','2018-05-13 10:13:07'),
	(22,1,'admin/auth/menu','POST','127.0.0.1','{\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\",\"_order\":\"[{\\\"id\\\":1},{\\\"id\\\":2,\\\"children\\\":[{\\\"id\\\":3},{\\\"id\\\":4},{\\\"id\\\":5},{\\\"id\\\":6},{\\\"id\\\":7}]},{\\\"id\\\":9},{\\\"id\\\":10},{\\\"id\\\":11}]\"}','2018-05-13 10:13:12','2018-05-13 10:13:12'),
	(23,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:13:13','2018-05-13 10:13:13'),
	(24,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:13:18','2018-05-13 10:13:18'),
	(25,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:13:22','2018-05-13 10:13:22'),
	(26,1,'admin','GET','127.0.0.1','[]','2018-05-13 10:13:29','2018-05-13 10:13:29'),
	(27,1,'admin/ninepic','GET','127.0.0.1','[]','2018-05-13 10:14:19','2018-05-13 10:14:19'),
	(28,1,'admin/ninepic','GET','127.0.0.1','[]','2018-05-13 10:20:22','2018-05-13 10:20:22'),
	(29,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:20:25','2018-05-13 10:20:25'),
	(30,1,'admin/user','GET','127.0.0.1','[]','2018-05-13 10:21:02','2018-05-13 10:21:02'),
	(31,1,'admin/user','GET','127.0.0.1','[]','2018-05-13 10:21:43','2018-05-13 10:21:43'),
	(32,1,'admin/user/create','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:21:45','2018-05-13 10:21:45'),
	(33,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:21:47','2018-05-13 10:21:47'),
	(34,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:22:05','2018-05-13 10:22:05'),
	(35,1,'admin/config','GET','127.0.0.1','[]','2018-05-13 10:27:04','2018-05-13 10:27:04'),
	(36,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:06','2018-05-13 10:27:06'),
	(37,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:09','2018-05-13 10:27:09'),
	(38,1,'admin/auth/menu/9/edit','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:11','2018-05-13 10:27:11'),
	(39,1,'admin/auth/menu/9','PUT','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u4e5d\\u56fe\",\"icon\":\"fa-adjust\",\"uri\":\"\\/ninepic\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/127.0.0.1:8000\\/admin\\/auth\\/menu\"}','2018-05-13 10:27:16','2018-05-13 10:27:16'),
	(40,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:27:17','2018-05-13 10:27:17'),
	(1612,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 10:01:05','2018-09-08 10:01:05'),
	(1613,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 10:01:33','2018-09-08 10:01:33'),
	(1614,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 10:03:12','2018-09-08 10:03:12'),
	(1615,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 10:04:16','2018-09-08 10:04:16'),
	(1616,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 10:07:35','2018-09-08 10:07:35'),
	(1617,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 10:07:38','2018-09-08 10:07:38'),
	(1618,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 10:07:39','2018-09-08 10:07:39'),
	(1619,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 10:07:40','2018-09-08 10:07:40'),
	(1620,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:08:19','2018-09-08 10:08:19'),
	(1621,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:09:05','2018-09-08 10:09:05'),
	(1622,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 10:14:42','2018-09-08 10:14:42'),
	(1623,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 10:14:43','2018-09-08 10:14:43'),
	(1624,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 10:14:44','2018-09-08 10:14:44'),
	(1625,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 10:14:44','2018-09-08 10:14:44'),
	(1626,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 10:14:47','2018-09-08 10:14:47'),
	(1627,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:39:06','2018-09-08 10:39:06'),
	(1628,1,'/admin/info/roles','GET','127.0.0.1','','2018-09-08 10:39:07','2018-09-08 10:39:07'),
	(1629,1,'/admin/info/permission','GET','127.0.0.1','','2018-09-08 10:39:08','2018-09-08 10:39:08'),
	(1630,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 10:39:09','2018-09-08 10:39:09'),
	(1631,1,'/admin/info/op','GET','127.0.0.1','','2018-09-08 10:39:10','2018-09-08 10:39:10'),
	(1632,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 10:39:13','2018-09-08 10:39:13'),
	(1633,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 10:39:14','2018-09-08 10:39:14'),
	(1634,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 10:39:15','2018-09-08 10:39:15'),
	(1635,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:40:56','2018-09-08 10:40:56'),
	(1636,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:44:50','2018-09-08 10:44:50'),
	(1637,1,'/admin/info/manager/new','GET','127.0.0.1','','2018-09-08 10:44:53','2018-09-08 10:44:53'),
	(1638,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:44:54','2018-09-08 10:44:54'),
	(1639,1,'/admin/info/manager/edit','GET','127.0.0.1','','2018-09-08 10:44:55','2018-09-08 10:44:55'),
	(1640,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 10:44:56','2018-09-08 10:44:56'),
	(1641,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 12:48:13','2018-09-08 12:48:13'),
	(1642,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 12:48:15','2018-09-08 12:48:15'),
	(1643,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"mksJmNJZ2hR3VIdPnfx5m3gme0GEDCIlBE7\"],\"city\":[\"North Gideonshire\"],\"gender\":[\"0\"],\"id\":[\"3632\"],\"ip\":[\"199.1.120.7\"],\"name\":[\"john\"],\"phone\":[\"1-739-973-\"]}','2018-09-08 12:48:22','2018-09-08 12:48:22'),
	(1644,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 12:48:26','2018-09-08 12:48:26'),
	(1645,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"tWRO1tbgRqvqvdxx0Sowl4zJ94iJVEjDQ3r\"],\"city\":[\"North Gideonshire\"],\"gender\":[\"1\"],\"id\":[\"3632\"],\"ip\":[\"199.1.120.7\"],\"name\":[\"john\"],\"phone\":[\"1-739-973-\"]}','2018-09-08 12:48:29','2018-09-08 12:48:29'),
	(1646,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 12:48:35','2018-09-08 12:48:35'),
	(1647,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"uXUv8eFk3r0TBzHmFHvGfSvjO0Cb2ZFO6KZ\"],\"city\":[\"Suzanneville\"],\"gender\":[\"2\"],\"id\":[\"3631\"],\"ip\":[\"242.42.173.68\"],\"name\":[\"iure\"],\"phone\":[\"446.892.99\"]}','2018-09-08 12:48:38','2018-09-08 12:48:38'),
	(1648,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 12:48:45','2018-09-08 12:48:45'),
	(1649,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 12:48:46','2018-09-08 12:48:46'),
	(1650,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 12:48:47','2018-09-08 12:48:47'),
	(1651,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 12:50:10','2018-09-08 12:50:10'),
	(1652,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 12:50:12','2018-09-08 12:50:12'),
	(1653,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 12:51:13','2018-09-08 12:51:13'),
	(1654,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 12:51:15','2018-09-08 12:51:15'),
	(1655,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"MeAKPqWvf90VwTwn3LmD81rTQk350d1OKY7\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+861534673873xx\"]}','2018-09-08 12:51:44','2018-09-08 12:51:44'),
	(1656,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 12:51:55','2018-09-08 12:51:55'),
	(1657,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"4zojGavjyejMTtTEyCGnJmgDAfWwE0bUAdz\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+861534673873xx\"]}','2018-09-08 12:52:57','2018-09-08 12:52:57'),
	(1658,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 12:58:44','2018-09-08 12:58:44'),
	(1659,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"F1M7rK6CnjcTtHYprZzPitisbdeG707scXd\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+861534673873xx\"]}','2018-09-08 12:59:20','2018-09-08 12:59:20'),
	(1660,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 13:01:09','2018-09-08 13:01:09'),
	(1661,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"yE1Ph6EC3i2GeWGFers0saQOf2LOFn5xiYR\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+861534673873xx\"]}','2018-09-08 13:01:12','2018-09-08 13:01:12'),
	(1662,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 13:04:32','2018-09-08 13:04:32'),
	(1663,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"Qb5m6bMmvEmkZ8zqmXbmsYwjxaLxZk91xm1\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+861534673873xx\"]}','2018-09-08 13:04:35','2018-09-08 13:04:35'),
	(1664,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 13:14:30','2018-09-08 13:14:30'),
	(1665,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 13:14:33','2018-09-08 13:14:33'),
	(1666,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"gWhxlWEXLT9KSsnLBjcMlGnMZUlDOt7S3zz\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"+8613929345678xx\"]}','2018-09-08 13:15:00','2018-09-08 13:15:00'),
	(1667,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:17:40','2018-09-08 15:17:40'),
	(1668,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"0aGFbBkmT5EWhce7V2LUXHbK9a7n8YH2kcc\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"ray\"],\"phone\":[\"12312312332132\"]}','2018-09-08 15:18:05','2018-09-08 15:18:05'),
	(1669,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"PByMV8gWAcZfsLgVz57K0zl2nUcjEg8vr3X\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"23231212132131\"]}','2018-09-08 15:19:12','2018-09-08 15:19:12'),
	(1670,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:19:34','2018-09-08 15:19:34'),
	(1671,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"WryznQrPjvAau43foZPYYPNcBB6dhumK0yI\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"23213131131131313\"]}','2018-09-08 15:19:53','2018-09-08 15:19:53'),
	(1672,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 15:26:14','2018-09-08 15:26:14'),
	(1673,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:26:17','2018-09-08 15:26:17'),
	(1674,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"zVJFxHjadSmlytEJ9zUeldF0tVnGuikr41c\"],\"city\":[\"GuangZhou\"],\"gender\":[\"1\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"+86 672345678923232\"]}','2018-09-08 15:27:15','2018-09-08 15:27:15'),
	(1675,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:29:40','2018-09-08 15:29:40'),
	(1676,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:29:41','2018-09-08 15:29:41'),
	(1677,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"6KtFxPTnXdwhcMok9cKRUttIKfvWvbJfSv7\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"+86121392934562323\"]}','2018-09-08 15:30:01','2018-09-08 15:30:01'),
	(1678,1,'/admin/info/user/new','GET','127.0.0.1','','2018-09-08 15:30:17','2018-09-08 15:30:17'),
	(1679,1,'/admin/new/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"ICLENgULcDRiO5wBuKvyZxBlsCZqHZrAxOf\"],\"city\":[\"GuangZhou\"],\"gender\":[\"2\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"+86123232\"]}','2018-09-08 15:30:40','2018-09-08 15:30:40'),
	(1680,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 15:30:50','2018-09-08 15:30:50'),
	(1681,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"hvKxEhuaXKHkwRp53jBOdQ0oD3bhdwPvv6L\"],\"city\":[\"GuangZhou\"],\"gender\":[\"1\"],\"id\":[\"3633\"],\"ip\":[\"127.0.0.1\"],\"name\":[\"Ray\"],\"phone\":[\"+86123232\"]}','2018-09-08 15:30:53','2018-09-08 15:30:53'),
	(1682,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 15:38:50','2018-09-08 15:38:50'),
	(1683,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"vGEEIqVADVUhAS2UhvMgiHx6gPP4dWQs4uT\"],\"city\":[\"South Lydiaberg\"],\"gender\":[\"1\"],\"id\":[\"3630\"],\"ip\":[\"180.64.77.168\"],\"name\":[\"libero\"],\"phone\":[\"0835552395\"]}','2018-09-08 15:38:55','2018-09-08 15:38:55'),
	(1684,1,'/admin/info/user/edit','GET','127.0.0.1','','2018-09-08 15:38:57','2018-09-08 15:38:57'),
	(1685,1,'/admin/edit/user','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/user?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"njxAROWWvBQjz4CfFBhqsDD13dmOFLSFPq1\"],\"city\":[\"East Careyborough\"],\"gender\":[\"8\"],\"id\":[\"3629\"],\"ip\":[\"251.55.8.91\"],\"name\":[\"cupiditate\"],\"phone\":[\"(763)447-2\"]}','2018-09-08 15:39:00','2018-09-08 15:39:00'),
	(1686,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 15:39:53','2018-09-08 15:39:53'),
	(1687,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 15:39:54','2018-09-08 15:39:54'),
	(1688,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 15:39:55','2018-09-08 15:39:55'),
	(1689,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 15:40:03','2018-09-08 15:40:03'),
	(1690,1,'/admin/menu/order','POST','127.0.0.1','','2018-09-08 15:40:17','2018-09-08 15:40:17'),
	(1691,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 15:40:18','2018-09-08 15:40:18'),
	(1692,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 15:40:22','2018-09-08 15:40:22'),
	(1693,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 15:43:12','2018-09-08 15:43:12'),
	(1694,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 15:43:14','2018-09-08 15:43:14'),
	(1695,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 15:44:10','2018-09-08 15:44:10'),
	(1696,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:00:36','2018-09-08 16:00:36'),
	(1697,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:02:13','2018-09-08 16:02:13'),
	(1698,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:04:23','2018-09-08 16:04:23'),
	(1699,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:07:09','2018-09-08 16:07:09'),
	(1700,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:08:07','2018-09-08 16:08:07'),
	(1701,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:15:52','2018-09-08 16:15:52'),
	(1702,1,'/admin/delete/user','POST','127.0.0.1','','2018-09-08 16:15:59','2018-09-08 16:15:59'),
	(1703,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:15:59','2018-09-08 16:15:59'),
	(1704,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:16:17','2018-09-08 16:16:17'),
	(1705,1,'/admin/delete/user','POST','127.0.0.1','','2018-09-08 16:16:21','2018-09-08 16:16:21'),
	(1706,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:16:21','2018-09-08 16:16:21'),
	(1707,1,'/admin/delete/user','POST','127.0.0.1','','2018-09-08 16:16:50','2018-09-08 16:16:50'),
	(1708,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:16:50','2018-09-08 16:16:50'),
	(1709,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:21:25','2018-09-08 16:21:25'),
	(1710,1,'/admin/delete/user','POST','127.0.0.1','','2018-09-08 16:21:30','2018-09-08 16:21:30'),
	(1711,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:21:30','2018-09-08 16:21:30'),
	(1712,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:21:33','2018-09-08 16:21:33'),
	(1713,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:21:46','2018-09-08 16:21:46'),
	(1714,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:37:56','2018-09-08 16:37:56'),
	(1715,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 16:38:21','2018-09-08 16:38:21'),
	(1716,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:38:28','2018-09-08 16:38:28'),
	(1717,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:44:38','2018-09-08 16:44:38'),
	(1718,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:44:52','2018-09-08 16:44:52'),
	(1719,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:45:58','2018-09-08 16:45:58'),
	(1720,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:50:57','2018-09-08 16:50:57'),
	(1721,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:51:02','2018-09-08 16:51:02'),
	(1722,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:04','2018-09-08 16:51:04'),
	(1723,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:51:05','2018-09-08 16:51:05'),
	(1724,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:07','2018-09-08 16:51:07'),
	(1725,1,'/admin/info/user','GET','127.0.0.1','','2018-09-08 16:51:09','2018-09-08 16:51:09'),
	(1726,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:12','2018-09-08 16:51:12'),
	(1727,1,'/admin/delete/posts','POST','127.0.0.1','','2018-09-08 16:51:29','2018-09-08 16:51:29'),
	(1728,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:29','2018-09-08 16:51:29'),
	(1729,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:35','2018-09-08 16:51:35'),
	(1730,1,'/admin/info/posts/new','GET','127.0.0.1','','2018-09-08 16:51:37','2018-09-08 16:51:37'),
	(1731,1,'/admin/info/posts','GET','127.0.0.1','','2018-09-08 16:51:38','2018-09-08 16:51:38'),
	(1732,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 16:51:41','2018-09-08 16:51:41'),
	(1733,1,'/admin/info/manager/new','GET','127.0.0.1','','2018-09-08 16:51:43','2018-09-08 16:51:43'),
	(1734,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 16:51:47','2018-09-08 16:51:47'),
	(1735,1,'/admin/info/manager/new','GET','127.0.0.1','','2018-09-08 16:51:50','2018-09-08 16:51:50'),
	(1736,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-08 16:51:53','2018-09-08 16:51:53'),
	(1737,1,'/admin/info/roles','GET','127.0.0.1','','2018-09-08 16:51:54','2018-09-08 16:51:54'),
	(1738,1,'/admin/info/roles/new','GET','127.0.0.1','','2018-09-08 16:51:56','2018-09-08 16:51:56'),
	(1739,1,'/admin/info/roles','GET','127.0.0.1','','2018-09-08 16:51:58','2018-09-08 16:51:58'),
	(1740,1,'/admin/info/permission','GET','127.0.0.1','','2018-09-08 16:51:59','2018-09-08 16:51:59'),
	(1741,1,'/admin/info/permission/new','GET','127.0.0.1','','2018-09-08 16:52:00','2018-09-08 16:52:00'),
	(1742,1,'/admin/info/permission','GET','127.0.0.1','','2018-09-08 16:52:02','2018-09-08 16:52:02'),
	(1743,1,'/admin/menu','GET','127.0.0.1','','2018-09-08 16:52:03','2018-09-08 16:52:03'),
	(1744,1,'/admin/info/op','GET','127.0.0.1','','2018-09-08 16:52:05','2018-09-08 16:52:05'),
	(1745,1,'/admin/info/op/new','GET','127.0.0.1','','2018-09-08 16:52:07','2018-09-08 16:52:07'),
	(1746,1,'/admin/info/op','GET','127.0.0.1','','2018-09-08 16:52:08','2018-09-08 16:52:08'),
	(1747,1,'/admin/info/authors','GET','127.0.0.1','','2018-09-08 16:52:39','2018-09-08 16:52:39'),
	(1748,1,'/admin/info/manager','GET','127.0.0.1','','2018-09-09 08:58:41','2018-09-09 08:58:41'),
	(1749,1,'/admin/info/manager/edit','GET','127.0.0.1','','2018-09-09 08:58:44','2018-09-09 08:58:44'),
	(1750,1,'/admin/edit/manager','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/manager?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"vzvUnkSlnkTFjqJHgBnnL43FJad4ariaU2v\"],\"avatar\":[\"\"],\"id\":[\"1\"],\"name\":[\"admin1\"],\"password\":[\"admin\"],\"permission_id[]\":[\"2\",\"3\"],\"role_id[]\":[\"1\"],\"username\":[\"admin\"]}','2018-09-09 08:58:48','2018-09-09 08:58:48'),
	(1751,1,'/admin/info/manager/edit','GET','127.0.0.1','','2018-09-09 08:58:51','2018-09-09 08:58:51'),
	(1752,1,'/admin/edit/manager','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/manager?page=1\\u0026pageSize=10\\u0026sort=id\\u0026sort_type=desc\"],\"_t\":[\"unmND5aAPCbP2ZYedhLQAKfMu31PZ0519bY\"],\"avatar\":[\"\"],\"id\":[\"1\"],\"name\":[\"admin\"],\"password\":[\"admin\"],\"permission_id[]\":[\"2\",\"3\"],\"role_id[]\":[\"1\"],\"username\":[\"admin\"]}','2018-09-09 08:58:55','2018-09-09 08:58:55'),
	(1753,1,'/admin/info/manager/new','GET','127.0.0.1','','2018-09-09 08:58:58','2018-09-09 08:58:58'),
	(1754,1,'/admin/new/manager','POST','127.0.0.1','{\"_previous_\":[\"/admin/info/manager?page=1\\u0026pageSize=10\\u0026sort=desc\\u0026sort_type=desc\"],\"_t\":[\"fwd6HsTaoNdejKxDu5EXvbSc3N9wYAYJSYj\"],\"name\":[\"admin234\"],\"password\":[\"123456\"],\"permission_id[]\":[\"1\"],\"role_id[]\":[\"1\"],\"username\":[\"admin234\"]}','2018-09-09 09:00:22','2018-09-09 09:00:22'),
	(1755,1,'/admin/info/manager/new','GET','127.0.0.1','','2018-09-09 09:05:46','2018-09-09 09:05:46');

/*!40000 ALTER TABLE `goadmin_operation_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_permissions`;

CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `created_at`, `updated_at`)
VALUES
	(1,'All permission','*','','*',NULL,NULL),
	(2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/',NULL,NULL),
	(3,'Login','auth.login','','/auth/login\r\n/auth/logout',NULL,NULL),
	(4,'User setting','auth.setting','GET,PUT','/auth/setting',NULL,NULL),
	(5,'Auth management','auth.management','','/auth/roles\r\n/auth/permissions\r\n/auth/menu\r\n/auth/logs',NULL,NULL);

/*!40000 ALTER TABLE `goadmin_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_menu`;

CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_role_menu` DISABLE KEYS */;

INSERT INTO `goadmin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`)
VALUES
	(1,2,NULL,NULL),
	(1,11,NULL,NULL);

/*!40000 ALTER TABLE `goadmin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_permissions`;

CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_role_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,1,NULL,NULL),
	(1,2,'2018-08-26 11:31:35','2018-08-26 11:31:35');

/*!40000 ALTER TABLE `goadmin_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_role_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_role_users`;

CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_role_users` WRITE;
/*!40000 ALTER TABLE `goadmin_role_users` DISABLE KEYS */;

INSERT INTO `goadmin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`)
VALUES
	(1,1,NULL,NULL),
	(1,2,'2018-09-09 09:25:24','2018-09-09 09:25:24'),
	(1,3,'2018-09-09 09:27:06','2018-09-09 09:27:06'),
	(1,4,'2018-09-09 09:29:15','2018-09-09 09:29:15'),
	(1,5,'2018-08-10 23:06:18','2018-08-10 23:06:18'),
	(1,6,'2018-08-10 23:08:59','2018-08-10 23:08:59'),
	(1,7,'2018-08-10 23:13:12','2018-08-10 23:13:12'),
	(1,9,'2018-08-10 23:21:12','2018-08-10 23:21:12'),
	(1,10,'2018-08-10 23:25:36','2018-08-10 23:25:36'),
	(1,11,'2018-08-10 23:34:36','2018-08-10 23:34:36'),
	(1,12,'2018-08-10 23:36:04','2018-08-10 23:36:04'),
	(1,13,'2018-08-10 23:39:35','2018-08-10 23:39:35'),
	(1,14,'2018-08-10 23:42:09','2018-08-10 23:42:09');

/*!40000 ALTER TABLE `goadmin_role_users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_roles
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_roles`;

CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_roles` WRITE;
/*!40000 ALTER TABLE `goadmin_roles` DISABLE KEYS */;

INSERT INTO `goadmin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,'Administrator','administrator','2018-05-13 10:00:33','2018-05-13 10:00:33');

/*!40000 ALTER TABLE `goadmin_roles` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_session
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_session`;

CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) DEFAULT NULL,
  `values` varchar(3000) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `goadmin_session` WRITE;
/*!40000 ALTER TABLE `goadmin_session` DISABLE KEYS */;

INSERT INTO `goadmin_session` (`id`, `sid`, `values`, `created_at`, `updated_at`)
VALUES
	(1,'eFVSeWZEVFBrb2dBc2xrTndKYkRLZnhGekRQbWd4SHI=','{\"AiGg1gD7DQXdvchMZ42J7XWjM28A0MOE5JRXMEBvtQEaUvQbjlTWGbn91sYj\":\"1\"}','2018-08-12 12:08:22','2018-08-12 12:08:22'),
	(2,'a3J1WGxZZ2drenhCT0J3bXFVWW9la3paUkVRSEJieVc=','{\"mMq3KQgQE9Uxv6QfmBLIHOvYCUJV7tMDlKuhzRF3L66WyGlNfCxyHQCi9jql\":\"1\"}','2018-08-12 12:51:19','2018-08-12 12:51:19'),
	(3,'cnB5RkZta1R6VXFtY2dEZ0FyWk9waVh5ZmZxREVadXc=','{\"HtHNQZA2iefhUvnhn6pefUWRsIJpK9LxHkeJp3mPFMdUNAvdh2KvicHgvnTo\":\"1\",\"tWVTV6e84I20IW4hWINMHulAFbrQwkUgpNPfX7DfbBwpghWFNX4kdX8jS26F\":\"1\"}','2018-08-12 15:08:31','2018-08-12 15:08:31'),
	(4,'Qk50ZmFxdmtQZkNSck5GYnR3aE96Y05jWlBHZ0pCTXk=','{\"FnFVU52D9TH07rwpxDG4lHy7x0leU1M7IlIs80r1avk1QgQVConxcrkSm4zu\":\"1\"}','2018-08-12 17:09:57','2018-08-12 17:09:57'),
	(5,'b2FIclBmamp1UW9sYmhFT0xiaVhLbGVSZ2tXeVVhaEE=','{\"JpeFKyAwKinQsqKKPM2cwXXPzp8v4nTpnUw9bdtAsju0tZzvd794c1T6SDDQ\":\"1\"}','2018-08-12 22:25:55','2018-08-12 22:25:55'),
	(6,'Q3dxVHVJanFseExUYlNJcmp5ZEJCS0taQW9iYlVIcHI=','{\"vUF1XCMREITRODUFzvaNH8hTYahc0haHbiAl6naN6s9HB8GDYgz2ESRc5rm8\":\"1\"}','2018-08-13 08:45:00','2018-08-13 08:45:00'),
	(7,'ZWRyZ3ZOR0FadFBFRVFNQ0NaZHRDcVRtdkFNTlBsblc=','{\"user_id\":\"1\"}','2018-08-13 09:15:00','2018-08-13 09:15:00'),
	(8,'anRsQ2pVSlhXT1FtaGVpdHJSb1dsYVlXUFd6WU9FZ3c=','{\"user_id\":\"1\"}','2018-08-14 08:38:19','2018-08-14 08:38:19'),
	(9,'ZldXalFpaWJ1Z0dZdlZMdXNBWklxcWhpQlRrZXBudU4=','{\"user_id\":\"1\"}','2018-08-17 22:47:25','2018-08-17 22:47:25'),
	(10,'d0dUcnl6U0txdVdQUktpTUdQdk9YQ3ZVTGd2enBYRkI=','{\"user_id\":\"1\"}','2018-08-18 10:48:47','2018-08-18 10:48:47'),
	(11,'QkZNVkZLRWl4eUNZY2loU3h6SlBrYVRndmlLaGpxTnY=','{\"user_id\":\"1\"}','2018-08-19 18:22:12','2018-08-19 18:22:12'),
	(12,'WmpkWmtlY3hoYURwTVd1eFV5dGVHT3ZiR2RKRUhobWE=','{\"user_id\":\"1\"}','2018-08-20 08:54:42','2018-08-20 08:54:42'),
	(13,'UHZpblN2UW1LeFJkdUR1WEVIQU1Tc1BTQkdqRXdOSnU=','{\"user_id\":\"1\"}','2018-08-20 20:19:13','2018-08-20 20:19:13'),
	(14,'ZVdZU1lrYlZHU2NFaUdlamdsU0hsdFlzSnVCTXpzekY=','{\"user_id\":\"1\"}','2018-08-21 09:10:18','2018-08-21 09:10:18'),
	(15,'bEFBZFNUTXlRaVVsUnVnQnBSTU5BemN3bkNyQUx6Y3I=','{\"user_id\":\"1\"}','2018-08-21 22:40:37','2018-08-21 22:40:37'),
	(16,'T2FOeU1vTlRVUHprRFFITHBiWlduakx3bFJxbkxtYkU=','{\"user_id\":\"1\"}','2018-08-26 01:00:46','2018-08-26 01:00:46'),
	(17,'WUJpWkZYSmdjU2FsTUNqblhjdUxKcmRXbm1RQmppdUo=','{\"user_id\":\"1\"}','2018-08-26 11:01:32','2018-08-26 11:01:32'),
	(18,'nzNQG5qe4apZ3y3','{\"user_id\":\"1\"}','2018-09-01 00:25:26','2018-09-01 00:25:26'),
	(19,'1I0errloiiveGvh','{\"user_id\":\"1\"}','2018-09-01 00:25:27','2018-09-01 00:25:27'),
	(20,'NZrmqHeLkppMNvP','{\"user_id\":\"1\"}','2018-09-01 17:58:45','2018-09-01 17:58:45'),
	(21,'n9UQS35V8MlJpKA','{\"user_id\":\"1\"}','2018-09-02 09:09:03','2018-09-02 09:09:03'),
	(22,'0bSvcdeM5hazzgv','{\"user_id\":\"1\"}','2018-09-02 20:51:04','2018-09-02 20:51:04'),
	(23,'OVZzGqQEKvOnFWk','{\"user_id\":\"1\"}','2018-09-03 22:13:21','2018-09-03 22:13:21'),
	(24,'YdJlvjmcaY2iDut','{\"user_id\":\"1\"}','2018-09-04 08:23:13','2018-09-04 08:23:13'),
	(25,'v7It8ATzvg1qTeW','{\"user_id\":\"1\"}','2018-09-05 09:55:03','2018-09-05 09:55:03'),
	(26,'zNqWLogLIfYLPu6','{\"user_id\":\"1\"}','2018-09-06 19:34:01','2018-09-06 19:34:01'),
	(27,'YKKWYftL2F72mR2','{\"user_id\":\"1\"}','2018-09-07 08:18:35','2018-09-07 08:18:35'),
	(28,'kLXTe4NtwNnHv6c','{\"user_id\":\"1\"}','2018-09-07 21:05:46','2018-09-07 21:05:46'),
	(29,'ZJE7SdSTXzmVrBG','{\"user_id\":\"1\"}','2018-09-08 08:34:38','2018-09-08 08:34:38'),
	(30,'YCq6joKjTi8DBpf','{\"user_id\":\"1\"}','2018-09-09 08:18:39','2018-09-09 08:18:39'),
	(31,'8h4eX4ycywang7H','{\"user_id\":\"1\"}','2018-09-09 08:19:43','2018-09-09 08:19:43'),
	(33,'AlbK4574BNOmYHQ','{\"user_id\":\"1\"}','2018-09-09 09:54:53','2018-09-09 09:54:53'),
	(34,'CrEbhokvzFztQ3o','{\"user_id\":\"1\"}','2018-09-09 21:15:53','2018-09-09 21:15:53');

/*!40000 ALTER TABLE `goadmin_session` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_user_permissions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_user_permissions`;

CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_user_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_user_permissions` DISABLE KEYS */;

INSERT INTO `goadmin_user_permissions` (`user_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
	(1,2,NULL,NULL),
	(1,3,NULL,NULL),
	(2,1,'2018-09-09 09:25:24','2018-09-09 09:25:24'),
	(3,1,'2018-09-09 09:27:06','2018-09-09 09:27:06'),
	(4,1,'2018-09-09 09:29:15','2018-09-09 09:29:15'),
	(5,1,'2018-08-10 23:06:18','2018-08-10 23:06:18'),
	(5,2,'2018-08-10 23:06:18','2018-08-10 23:06:18'),
	(6,1,'2018-08-10 23:08:59','2018-08-10 23:08:59'),
	(6,2,'2018-08-10 23:08:59','2018-08-10 23:08:59'),
	(7,1,'2018-08-10 23:13:12','2018-08-10 23:13:12'),
	(9,1,'2018-08-10 23:21:12','2018-08-10 23:21:12'),
	(9,2,'2018-08-10 23:21:12','2018-08-10 23:21:12'),
	(10,1,'2018-08-10 23:25:36','2018-08-10 23:25:36'),
	(11,1,'2018-08-10 23:34:36','2018-08-10 23:34:36'),
	(12,1,'2018-08-10 23:36:04','2018-08-10 23:36:04'),
	(13,1,'2018-08-10 23:39:35','2018-08-10 23:39:35'),
	(14,1,'2018-08-10 23:42:09','2018-08-10 23:42:09');

/*!40000 ALTER TABLE `goadmin_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_users`;

CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(190) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(80) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_users` WRITE;
/*!40000 ALTER TABLE `goadmin_users` DISABLE KEYS */;

INSERT INTO `goadmin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`)
VALUES
	(1,'admin','$2a$10$FgN6YzUL1UZ4IjgDj/Fb5uaKb6zRXISTkM7/vonco1RxpvxMIrzdS','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2018-05-13 10:00:33','2018-05-13 10:00:33');

/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
