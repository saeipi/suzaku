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
DROP TABLE IF EXISTS `group_requests`;
CREATE TABLE `group_requests` (
  `user_id` varchar(40) DEFAULT '' COMMENT '事件ID',
  `group_id` varchar(40) DEFAULT '' COMMENT '发起人ID',
  `handle_user_id` varchar(40) DEFAULT '' COMMENT '处理人ID',
  `handle_result` tinyint(1) DEFAULT '0' COMMENT '结果',
  `handle_msg` varchar(255) DEFAULT '' COMMENT '处理消息',
  `handled_ts` bigint DEFAULT '0',
  `req_msg` varchar(255) DEFAULT '' COMMENT '添加好友消息',
  `req_ts` bigint DEFAULT '0' COMMENT '请求时间',
  `req_source` int DEFAULT '0' COMMENT '来源',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`user_id`,`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
