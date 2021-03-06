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
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `group_id` varchar(40) DEFAULT '',
  `group_name` varchar(60) DEFAULT '' COMMENT '名称',
  `avatar_url` varchar(255) DEFAULT '' COMMENT '头像',
  `notification` varchar(255) DEFAULT '' COMMENT '通知',
  `introduction` varchar(255) DEFAULT '' COMMENT '介绍',
  `creator_user_id` varchar(40) DEFAULT '' COMMENT '创建者ID',
  `group_type` int DEFAULT '0',
  `status` int DEFAULT '0',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`group_id`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
