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

 Date: 08/05/2022 16:46:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for group_avatars
-- ----------------------------
DROP TABLE IF EXISTS `group_avatars`;
CREATE TABLE `group_avatars` (
  `group_id` varchar(40) CHARACTER SET utf8 NOT NULL,
  `avatar_url` varchar(255) DEFAULT '' COMMENT '小图',
  `avatar_url_middle` varchar(255) DEFAULT '' COMMENT '中图',
  `avatar_url_big` varchar(255) DEFAULT '' COMMENT '大图',
  `updated_ts` bigint DEFAULT '0',
  PRIMARY KEY (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
