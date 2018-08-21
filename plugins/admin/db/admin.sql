# ************************************************************
# Sequel Pro SQL dump
# Version 4468
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.19)
# Database: godmin
# Generation Time: 2018-08-05 03:57:45 +0000
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;

INSERT INTO `goadmin_menu` (`id`, `parent_id`, `order`, `title`, `icon`, `uri`, `created_at`, `updated_at`)
VALUES
	(2,0,2,'管理','fa-tasks','',NULL,NULL),
	(3,2,3,'管理员管理','fa-users','/info/manager',NULL,NULL),
	(4,2,4,'角色管理','fa-user','/info/roles',NULL,NULL),
	(5,2,5,'权限管理','fa-ban','/info/permission',NULL,NULL),
	(6,2,6,'菜单管理','fa-bars','/menu',NULL,NULL),
	(7,2,7,'操作日志','fa-history','/info/op',NULL,NULL),
	(12,0,2,'用户表','fa-tasks','/info/user',NULL,NULL),
	(13,0,1,'仪表盘','fa-tasks','/','2018-08-03 15:24:42',NULL);

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
	(1,1,'admin','GET','127.0.0.1','[]','2018-05-13 10:06:51','2018-05-13 10:06:51'),
	(2,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:06:57','2018-05-13 10:06:57'),
	(3,1,'admin/auth/users','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:07:00','2018-05-13 10:07:00'),
	(4,1,'admin/auth/roles','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:08:54','2018-05-13 10:08:54'),
	(5,1,'admin/auth/permissions','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:08:56','2018-05-13 10:08:56'),
	(6,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:08:58','2018-05-13 10:08:58'),
	(7,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:09:02','2018-05-13 10:09:02'),
	(8,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:09:12','2018-05-13 10:09:12'),
	(9,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:11:41','2018-05-13 10:11:41'),
	(10,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:11:43','2018-05-13 10:11:43'),
	(11,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:11:49','2018-05-13 10:11:49'),
	(12,1,'admin/auth/menu','POST','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u4e5d\\u56fe\",\"icon\":\"fa-bars\",\"uri\":\"\\/admin\\/ninepic\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\"}','2018-05-13 10:12:14','2018-05-13 10:12:14'),
	(13,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:12:15','2018-05-13 10:12:15'),
	(14,1,'admin/auth/menu/8','DELETE','127.0.0.1','{\"_method\":\"delete\",\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\"}','2018-05-13 10:12:21','2018-05-13 10:12:21'),
	(15,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:12:21','2018-05-13 10:12:21'),
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
	(41,1,'admin/auth/menu/10/edit','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:20','2018-05-13 10:27:20'),
	(42,1,'admin/auth/menu/10','PUT','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u7528\\u6237\\u7ba1\\u7406\",\"icon\":\"fa-user\",\"uri\":\"\\/user\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/127.0.0.1:8000\\/admin\\/auth\\/menu\"}','2018-05-13 10:27:33','2018-05-13 10:27:33'),
	(43,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:27:33','2018-05-13 10:27:33'),
	(44,1,'admin/auth/menu/11/edit','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:36','2018-05-13 10:27:36'),
	(45,1,'admin/auth/menu/11','PUT','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u914d\\u7f6e\\u7ba1\\u7406\",\"icon\":\"fa-database\",\"uri\":\"\\/config\",\"roles\":[\"1\",null],\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/127.0.0.1:8000\\/admin\\/auth\\/menu\"}','2018-05-13 10:27:48','2018-05-13 10:27:48'),
	(46,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:27:48','2018-05-13 10:27:48'),
	(47,1,'admin/auth/menu','POST','127.0.0.1','{\"_token\":\"EgfqSyEZlex5p3uIg30VdX1YQEzraji1hWuarurC\",\"_order\":\"[{\\\"id\\\":1},{\\\"id\\\":2,\\\"children\\\":[{\\\"id\\\":3},{\\\"id\\\":4},{\\\"id\\\":5},{\\\"id\\\":6},{\\\"id\\\":7}]},{\\\"id\\\":9},{\\\"id\\\":10},{\\\"id\\\":11}]\"}','2018-05-13 10:27:50','2018-05-13 10:27:50'),
	(48,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:27:51','2018-05-13 10:27:51'),
	(49,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-05-13 10:27:53','2018-05-13 10:27:53'),
	(50,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:05','2018-05-13 10:28:05'),
	(51,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:14','2018-05-13 10:28:14'),
	(52,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:18','2018-05-13 10:28:18'),
	(53,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:20','2018-05-13 10:28:20'),
	(54,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:21','2018-05-13 10:28:21'),
	(55,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:24','2018-05-13 10:28:24'),
	(56,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:26','2018-05-13 10:28:26'),
	(57,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:27','2018-05-13 10:28:27'),
	(58,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:28','2018-05-13 10:28:28'),
	(59,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:29','2018-05-13 10:28:29'),
	(60,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:48','2018-05-13 10:28:48'),
	(61,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:49','2018-05-13 10:28:49'),
	(62,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:50','2018-05-13 10:28:50'),
	(63,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:51','2018-05-13 10:28:51'),
	(64,1,'admin/ninepic/create','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:56','2018-05-13 10:28:56'),
	(65,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-13 10:28:59','2018-05-13 10:28:59'),
	(66,1,'admin','GET','127.0.0.1','[]','2018-05-14 12:16:10','2018-05-14 12:16:10'),
	(67,1,'admin/auth/users','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 12:16:13','2018-05-14 12:16:13'),
	(68,1,'admin','GET','127.0.0.1','[]','2018-05-14 12:22:01','2018-05-14 12:22:01'),
	(69,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:17:03','2018-05-14 14:17:03'),
	(70,1,'admin/user','GET','127.0.0.1','[]','2018-05-14 14:17:07','2018-05-14 14:17:07'),
	(71,1,'admin/user','GET','127.0.0.1','[]','2018-05-14 14:17:10','2018-05-14 14:17:10'),
	(72,1,'admin/user','GET','127.0.0.1','[]','2018-05-14 14:17:12','2018-05-14 14:17:12'),
	(73,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:17:19','2018-05-14 14:17:19'),
	(74,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:17:21','2018-05-14 14:17:21'),
	(75,1,'admin/ninepic','GET','127.0.0.1','[]','2018-05-14 14:17:36','2018-05-14 14:17:36'),
	(76,1,'admin/config','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:17:58','2018-05-14 14:17:58'),
	(77,1,'admin/user','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:17:59','2018-05-14 14:17:59'),
	(78,1,'admin/ninepic','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:18:01','2018-05-14 14:18:01'),
	(79,1,'admin/auth/users','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-14 14:18:06','2018-05-14 14:18:06'),
	(80,1,'admin/auth/users','GET','127.0.0.1','[]','2018-05-14 15:05:30','2018-05-14 15:05:30'),
	(81,1,'admin/auth/users','GET','127.0.0.1','[]','2018-05-15 09:08:01','2018-05-15 09:08:01'),
	(82,1,'admin/auth/logout','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-05-15 09:08:06','2018-05-15 09:08:06'),
	(83,1,'admin','GET','127.0.0.1','[]','2018-06-13 08:39:56','2018-06-13 08:39:56'),
	(84,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:40:30','2018-06-13 08:40:30'),
	(85,1,'admin/auth/menu/9','DELETE','127.0.0.1','{\"_method\":\"delete\",\"_token\":\"qRcjb0uivYgFkpmClfg9VNK0W1EQeUqLhQXTS9yX\"}','2018-06-13 08:40:37','2018-06-13 08:40:37'),
	(86,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:40:37','2018-06-13 08:40:37'),
	(87,1,'admin/auth/menu/10','DELETE','127.0.0.1','{\"_method\":\"delete\",\"_token\":\"qRcjb0uivYgFkpmClfg9VNK0W1EQeUqLhQXTS9yX\"}','2018-06-13 08:40:40','2018-06-13 08:40:40'),
	(88,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:40:41','2018-06-13 08:40:41'),
	(89,1,'admin/auth/menu/11/edit','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:40:43','2018-06-13 08:40:43'),
	(90,1,'admin/auth/menu/11','PUT','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u68c0\\u6d4b\\u7f51\\u7ad9\",\"icon\":\"fa-database\",\"uri\":\"\\/web\",\"roles\":[\"1\",null],\"_token\":\"qRcjb0uivYgFkpmClfg9VNK0W1EQeUqLhQXTS9yX\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/127.0.0.1:8000\\/admin\\/auth\\/menu\"}','2018-06-13 08:41:01','2018-06-13 08:41:01'),
	(91,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-06-13 08:41:01','2018-06-13 08:41:01'),
	(92,1,'admin/auth/menu/11/edit','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:41:45','2018-06-13 08:41:45'),
	(93,1,'admin/auth/menu/11','PUT','127.0.0.1','{\"parent_id\":\"0\",\"title\":\"\\u68c0\\u6d4b\\u7f51\\u7ad9\",\"icon\":\"fa-database\",\"uri\":\"\\/\",\"roles\":[\"1\",null],\"_token\":\"qRcjb0uivYgFkpmClfg9VNK0W1EQeUqLhQXTS9yX\",\"_method\":\"PUT\",\"_previous_\":\"http:\\/\\/127.0.0.1:8000\\/admin\\/auth\\/menu\"}','2018-06-13 08:41:50','2018-06-13 08:41:50'),
	(94,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-06-13 08:41:50','2018-06-13 08:41:50'),
	(95,1,'admin/auth/menu/1','DELETE','127.0.0.1','{\"_method\":\"delete\",\"_token\":\"qRcjb0uivYgFkpmClfg9VNK0W1EQeUqLhQXTS9yX\"}','2018-06-13 08:41:58','2018-06-13 08:41:58'),
	(96,1,'admin/auth/menu','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:41:58','2018-06-13 08:41:58'),
	(97,1,'admin/auth/menu','GET','127.0.0.1','[]','2018-06-13 08:42:00','2018-06-13 08:42:00'),
	(98,1,'admin','GET','127.0.0.1','{\"_pjax\":\"#pjax-container\"}','2018-06-13 08:42:03','2018-06-13 08:42:03'),
	(99,1,'admin','GET','127.0.0.1','[]','2018-06-13 08:42:20','2018-06-13 08:42:20');

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
	(1,1,NULL,NULL);

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
	(1,1,NULL,NULL);

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
	(1,3,NULL,NULL);

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
	(1,'admin','$2a$10$YDoHIAPcGpa3/Pm0f5Q/HeAlhOaRUgL.eyF8Ne/Mc1dp9esEQEV5e','admin','/uploads/R5JLbqE6kmAl4qFMQdSgPMjaBfQfj9KFsQ64WvbboKF5X5jmT6','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2018-05-13 10:00:33','2018-05-13 10:00:33');

/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table goadmin_session
# ------------------------------------------------------------

DROP TABLE IF EXISTS `goadmin_session`;

CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(100) DEFAULT NULL,
  `values` varchar(3000) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
