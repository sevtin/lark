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
// 复合索引支持唯一性约束
db.messages.createIndex({chat_type:1,chat_id:1,seq_id:-1},{unique:true})
db.messages.createIndex({srv_msg_id:1,sender_id:1,seq_id:-1})
