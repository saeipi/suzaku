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
DROP TABLE IF EXISTS `friends`;
CREATE TABLE `friends` (
  `owner_user_id` varchar(40) DEFAULT '' COMMENT '添加好友发起者ID',
  `friend_user_id` varchar(40) DEFAULT '' COMMENT '好友ID',
  `operator_user_id` varchar(40) DEFAULT '' COMMENT '处理人ID',
  `session_id` char(40) CHARACTER SET utf8 DEFAULT '' COMMENT '会话ID',
  `source` tinyint(1) DEFAULT '0' COMMENT '添加源',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_ts` bigint DEFAULT '0',
  `updated_ts` bigint DEFAULT '0',
  `deleted_ts` bigint DEFAULT '0',
  PRIMARY KEY (`owner_user_id`,`friend_user_id`),
  KEY `idx_ownerUserId_deletedTs` (`owner_user_id`,`deleted_ts`),
  KEY `idx_ownerUserId_friendUserId` (`owner_user_id`,`friend_user_id`,`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
