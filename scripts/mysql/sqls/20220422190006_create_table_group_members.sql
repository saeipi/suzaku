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

 Date: 05/05/2022 19:40:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for group_members
-- ----------------------------
DROP TABLE IF EXISTS `group_members`;
CREATE TABLE `group_members` (
  `group_id` varchar(64) NOT NULL COMMENT '群ID',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `nickname` varchar(255) DEFAULT '' COMMENT '在群中的昵称',
  `user_avatar_url` varchar(255) DEFAULT '' COMMENT '在群中的头像',
  `role_level` int DEFAULT '0' COMMENT '角色等级',
  `join_time` bigint DEFAULT '0' COMMENT '加入时间',
  `join_source` int DEFAULT '0' COMMENT '来源',
  `operator_user_id` varchar(64) DEFAULT '' COMMENT '操作员',
  `mute_end_time` bigint DEFAULT '0' COMMENT '禁言结束时间',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`group_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
