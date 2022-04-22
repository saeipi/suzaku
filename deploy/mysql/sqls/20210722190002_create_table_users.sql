/*
 Navicat Premium Data Transfer

 Source Server         : ayumi
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : suzaku

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 22/04/2022 09:41:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(24) DEFAULT '' COMMENT '用户ID',
  `mobile` varchar(32) DEFAULT '' COMMENT '手机',
  `platform_id` tinyint(1) DEFAULT '0' COMMENT '平台',
  `gender` tinyint(1) DEFAULT '0' COMMENT '性别',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_userId` (`user_id`) USING BTREE,
  KEY `idx_ mobile` (`mobile`) USING BTREE,
  KEY `idx_platformId` (`platform_id`) USING BTREE,
  KEY `idx_gender` (`gender`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
