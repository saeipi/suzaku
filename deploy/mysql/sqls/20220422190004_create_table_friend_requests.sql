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
DROP TABLE IF EXISTS `friend_requests`;
CREATE TABLE `friend_requests` (
  `req_id` varchar(40) DEFAULT '' COMMENT '事件ID',
  `from_user_id` varchar(40) DEFAULT '' COMMENT '发起人ID',
  `to_user_id` varchar(40) DEFAULT '' COMMENT '目标人ID',
  `operator_user_id` varchar(40) DEFAULT '' COMMENT '处理人ID',
  `handle_result` tinyint(1) DEFAULT '0' COMMENT '结果',
  `req_msg` varchar(255) DEFAULT '' COMMENT '添加好友消息',
  `handle_msg` varchar(255) DEFAULT '' COMMENT '处理消息',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_at` datetime DEFAULT NULL,
  `handle_at` datetime DEFAULT NULL COMMENT '处理时间',
  PRIMARY KEY (`req_id`),
  KEY `idx_fromUserId` (`from_user_id`),
  KEY `idx_toUserId` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
