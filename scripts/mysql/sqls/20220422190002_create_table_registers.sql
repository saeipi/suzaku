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
DROP TABLE IF EXISTS `registers`;
CREATE TABLE `registers` (
  `user_id` varchar(40) DEFAULT '' COMMENT '用户ID 系统生成',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  PRIMARY KEY (`user_id`),
  KEY `idx_userId_password` (`user_id`,`password`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
