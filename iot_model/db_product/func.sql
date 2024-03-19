/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 5.7.36-log : Database - iot_product
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`iot_product` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `iot_product`;

/* Function  structure for function  `getProductTypeFullName` */

/*!50003 DROP FUNCTION IF EXISTS `getProductTypeFullName` */;
DELIMITER $$

/*!50003 CREATE DEFINER=`proxy`@`%` FUNCTION `getProductTypeFullName`(`id` bigint(20)) RETURNS varchar(512) CHARSET utf8mb4
    READS SQL DATA
BEGIN
	DECLARE fullName VARCHAR(1000);
	DECLARE parentId BIGINT(20);
	DECLARE parentName VARCHAR(1000);

	SET fullName = (SELECT a.name FROM t_pm_product_type a WHERE a.id = id);
	SET parentId = (SELECT a.parent_id FROM t_pm_product_type a WHERE a.id = id);

	WHILE parentId <> 0 DO
	SET parentName = (SELECT a.name FROM t_pm_product_type a WHERE a.id = parentId);
	SET fullName = CONCAT(parentName,"/",fullName);
	SET parentId = (SELECT a.parent_id FROM t_pm_product_type a WHERE a.id = parentId);
	END WHILE;
	RETURN fullName;
END */$$
DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
