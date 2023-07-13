/*
 Navicat Premium Data Transfer

 Source Server         : lark-lbs
 Source Server Type    : MongoDB
 Source Server Version : 50009
 Source Host           : localhost:27017
 Source Schema         : lark

 Target Server Type    : MongoDB
 Target Server Version : 50009
 File Encoding         : 65001

 Date: 12/07/2023 22:45:49
*/


// ----------------------------
// Collection structure for user_locations
// ----------------------------
db.getCollection("user_locations").drop();
db.createCollection("user_locations");
db.user_locations.ensureIndex({ location: "2dsphere"});
db.user_locations.createIndex( { "uid": 1 }, { unique: true } )
db.user_locations.createIndex({"gender":1})
db.user_locations.createIndex({"birth_ts":1})
db.user_locations.createIndex({"online_ts":1})

