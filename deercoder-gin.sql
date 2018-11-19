/*
 Navicat MySQL Data Transfer

 Source Server         : lu
 Source Server Type    : MySQL
 Source Server Version : 50724
 Source Host           : localhost:3306
 Source Schema         : deercoder-gin

 Target Server Type    : MySQL
 Target Server Version : 50724
 File Encoding         : 65001

 Date: 19/11/2018 14:25:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'username',
  `createtime` datetime(0) NULL DEFAULT '2018-08-14 11:47:53',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '梦', '2018-08-15 00:00:00');
INSERT INTO `user` VALUES (4, '梦', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (5, '梦', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (6, '梦', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (7, '梦', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (8, '梦', '2018-08-08 00:00:00');

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `userinfo` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `updatetime` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES (1, 1, '1', '2018-09-05 20:39:39');
INSERT INTO `userinfo` VALUES (2, 1, '哈j', '2018-09-06 11:15:24');

SET FOREIGN_KEY_CHECKS = 1;
