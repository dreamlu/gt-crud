/*
 Navicat MySQL Data Transfer

 Source Server         : lu
 Source Server Type    : MySQL
 Source Server Version : 50723
 Source Host           : localhost:3306
 Source Schema         : gindemo

 Target Server Type    : MySQL
 Target Server Version : 50723
 File Encoding         : 65001

 Date: 25/08/2018 21:28:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `userpassword` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `createtime` datetime(0) NULL DEFAULT '2018-08-14 11:47:53',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '梦鹿', '', '2018-08-09 00:00:00');
INSERT INTO `user` VALUES (3, '', '', '2018-08-14 11:47:53');
INSERT INTO `user` VALUES (4, '梦', '', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (5, '梦', '', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (6, '梦', '', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (7, '梦', '', '2018-08-08 00:00:00');
INSERT INTO `user` VALUES (8, '梦', '', '2018-08-08 00:00:00');

SET FOREIGN_KEY_CHECKS = 1;
