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
  `user_id` varchar(40) DEFAULT '' COMMENT '用户ID 系统生成',
  `szk_id` varchar(40) DEFAULT '' COMMENT '账户ID 用户设置',
  `nickname` varchar(60) DEFAULT '' COMMENT '昵称',
  `gender` tinyint(1) DEFAULT '0' COMMENT '性别',
  `birth` int DEFAULT '0' COMMENT '生日 时间戳',
  `email` varchar(64) DEFAULT '' COMMENT 'Email',
  `mobile` varchar(32) DEFAULT '' COMMENT '手机号',
  `platform_id` tinyint(1) DEFAULT '0' COMMENT '平台',
  `avatar_url` varchar(255) DEFAULT '' COMMENT '头像',
  `country` varchar(128) DEFAULT '' COMMENT '国家',
  `city` varchar(128) DEFAULT '' COMMENT '城市',
  `ex` varchar(255) DEFAULT '' COMMENT '扩展字段',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_szkId` (`szk_id`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_platformId` (`platform_id`),
  KEY `idx_gender` (`gender`),
  KEY `id_gender_city` (`gender`,`city`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;