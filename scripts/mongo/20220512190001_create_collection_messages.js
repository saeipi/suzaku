/*
 Navicat Premium Data Transfer

 Source Server         : saeipi
 Source Server Type    : MongoDB
 Source Server Version : 50008
 Source Host           : localhost:27017
 Source Schema         : suzaku

 Target Server Type    : MongoDB
 Target Server Version : 50008
 File Encoding         : 65001

 Date: 12/05/2022 20:43:20
*/


// ----------------------------
// Collection structure for messages
// ----------------------------
db.getCollection("messages").drop();
db.createCollection("messages");

db.messages.createIndex({"session_type":1,"session_id":1,"seq":-1})