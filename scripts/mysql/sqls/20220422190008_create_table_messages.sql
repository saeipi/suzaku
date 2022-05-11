/*
 Navicat Premium Data Transfer

 Source Server         : ayumi
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : 127.0.0.1:3306
 Source Schema         : sdb

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 07/05/2022 16:55:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `server_msg_id` char(40) CHARACTER SET utf8 NOT NULL COMMENT '服务端生成',
  `client_msg_id` char(40) CHARACTER SET utf8 DEFAULT '' COMMENT '客户端生成',
  `send_id` char(40) CHARACTER SET utf8 DEFAULT '' COMMENT '发送人ID',
  `recv_id` char(40) CHARACTER SET utf8 DEFAULT '' COMMENT '接收人ID 或 群ID',
  `sender_platform_id` tinyint DEFAULT '0' COMMENT '发送人平台ID',
  `sender_nickname` varchar(60) CHARACTER SET utf8 DEFAULT '',
  `sender_avatar_url` varchar(255) CHARACTER SET utf8 DEFAULT '',
  `session_id` char(40) CHARACTER SET utf8 DEFAULT '' COMMENT '单例:会话ID,群聊:群ID',
  `session_type` tinyint DEFAULT '0' COMMENT '1:单聊 2:群聊',
  `msg_from` int DEFAULT '0' COMMENT '100:用户消息 200:系统消息',
  `content_type` int DEFAULT '0',
  `content` varchar(3000) CHARACTER SET utf8 DEFAULT '',
  `status` tinyint DEFAULT '0',
  `send_ts` bigint DEFAULT '0' COMMENT '消息发送的具体时间(毫秒)',
  `created_ts` bigint DEFAULT '0' COMMENT '创建消息的时间，在send_ts之前',
  `ex` varchar(255) CHARACTER SET utf8 DEFAULT '',
  PRIMARY KEY (`server_msg_id`),
  KEY `idx_sessionType_sessionId_createdTs` (`session_type`,`session_id`,`created_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
