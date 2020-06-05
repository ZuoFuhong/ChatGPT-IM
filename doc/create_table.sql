/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : 127.0.0.1:3306
 Source Schema         : gim

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 05/06/2020 19:25:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `name` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `private_key` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '私钥',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='app';

-- ----------------------------
-- Records of app
-- ----------------------------
BEGIN;
INSERT INTO `app` VALUES (1, '测试', '-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQDcGsUIIAINHfRTdMmgGwLrjzfMNSrtgIf4EGsNaYwmC1GjF/bM\nh0Mcm10oLhNrKNYCTTQVGGIxuc5heKd1gOzb7bdTnCDPPZ7oV7p1B9Pud+6zPaco\nqDz2M24vHFWYY2FbIIJh8fHhKcfXNXOLovdVBE7Zy682X1+R1lRK8D+vmQIDAQAB\nAoGAeWAZvz1HZExca5k/hpbeqV+0+VtobMgwMs96+U53BpO/VRzl8Cu3CpNyb7HY\n64L9YQ+J5QgpPhqkgIO0dMu/0RIXsmhvr2gcxmKObcqT3JQ6S4rjHTln49I2sYTz\n7JEH4TcplKjSjHyq5MhHfA+CV2/AB2BO6G8limu7SheXuvECQQDwOpZrZDeTOOBk\nz1vercawd+J9ll/FZYttnrWYTI1sSF1sNfZ7dUXPyYPQFZ0LQ1bhZGmWBZ6a6wd9\nR+PKlmJvAkEA6o32c/WEXxW2zeh18sOO4wqUiBYq3L3hFObhcsUAY8jfykQefW8q\nyPuuL02jLIajFWd0itjvIrzWnVmoUuXydwJAXGLrvllIVkIlah+lATprkypH3Gyc\nYFnxCTNkOzIVoXMjGp6WMFylgIfLPZdSUiaPnxby1FNM7987fh7Lp/m12QJAK9iL\n2JNtwkSR3p305oOuAz0oFORn8MnB+KFMRaMT9pNHWk0vke0lB1sc7ZTKyvkEJW0o\neQgic9DvIYzwDUcU8wJAIkKROzuzLi9AvLnLUrSdI6998lmeYO9x7pwZPukz3era\nzncjRK3pbVkv0KrKfczuJiRlZ7dUzVO0b6QJr8TRAA==\n-----END RSA PRIVATE KEY-----', '2019-10-15 16:49:39', '2019-10-15 16:49:39');
COMMIT;

-- ----------------------------
-- Table structure for device
-- ----------------------------
DROP TABLE IF EXISTS `device`;
CREATE TABLE `device` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `device_id` bigint(20) NOT NULL COMMENT '设备id',
  `app_id` bigint(20) unsigned NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '账户id',
  `type` tinyint(4) NOT NULL COMMENT '设备类型,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web',
  `brand` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '手机厂商',
  `model` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '机型',
  `system_version` varchar(10) COLLATE utf8mb4_bin NOT NULL COMMENT '系统版本',
  `sdk_version` varchar(10) COLLATE utf8mb4_bin NOT NULL COMMENT 'app版本',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '在线状态，0：离线；1：在线',
  `conn_addr` varchar(25) COLLATE utf8mb4_bin NOT NULL COMMENT '连接层服务器地址',
  `conn_fd` bigint(20) NOT NULL COMMENT 'TCP连接对应的文件描述符',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_device_id` (`device_id`) USING BTREE,
  KEY `idx_app_id_user_id` (`app_id`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备';

-- ----------------------------
-- Records of device
-- ----------------------------
BEGIN;
INSERT INTO `device` VALUES (1, 56805130707664896, 1, 54146910402904064, 5, 'chrome', 'mac', '0.0.1', '0.0.1', 0, '', 0, '2020-06-05 18:03:18', '2020-06-05 18:03:18');
INSERT INTO `device` VALUES (2, 56805286664470528, 1, 54146910402904065, 5, 'chrome', 'mac', '0.0.1', '0.0.1', 0, '', 0, '2020-06-05 18:03:55', '2020-06-05 18:03:55');
INSERT INTO `device` VALUES (3, 56805315022159872, 1, 54146910402904066, 5, 'chrome', 'mac', '0.0.1', '0.0.1', 0, '', 0, '2020-06-05 18:04:02', '2020-06-05 18:04:02');
COMMIT;

-- ----------------------------
-- Table structure for device_ack
-- ----------------------------
DROP TABLE IF EXISTS `device_ack`;
CREATE TABLE `device_ack` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `device_id` bigint(20) unsigned NOT NULL COMMENT '设备id',
  `ack` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '收到消息确认号',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_device_id` (`device_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='设备消息同步序列号';

-- ----------------------------
-- Records of device_ack
-- ----------------------------
BEGIN;
INSERT INTO `device_ack` VALUES (1, 56805130707664896, 0, '2020-06-05 18:03:18', '2020-06-05 18:03:18');
INSERT INTO `device_ack` VALUES (2, 56805286664470528, 0, '2020-06-05 18:03:55', '2020-06-05 18:03:55');
INSERT INTO `device_ack` VALUES (3, 56805315022159872, 0, '2020-06-05 18:04:02', '2020-06-05 18:04:02');
COMMIT;

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) unsigned NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `friend_id` bigint(20) unsigned NOT NULL COMMENT '朋友ID',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_user_id_friend_id` (`app_id`,`user_id`,`friend_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='好友表';

-- ----------------------------
-- Records of friend
-- ----------------------------
BEGIN;
INSERT INTO `friend` VALUES (1, 1, 54146910402904064, 54146910402904065, '2020-06-05 17:41:43', '2020-06-05 17:41:43');
INSERT INTO `friend` VALUES (2, 1, 54146910402904064, 54146910402904066, '2020-06-05 17:42:16', '2020-06-05 17:42:16');
COMMIT;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) NOT NULL COMMENT 'app_id',
  `group_id` bigint(20) NOT NULL COMMENT '群组id',
  `name` varchar(50) COLLATE utf8mb4_bin NOT NULL COMMENT '群组名称',
  `avatar_url` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '群头像',
  `introduction` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '群组简介',
  `user_num` int(11) NOT NULL DEFAULT '0' COMMENT '群组人数',
  `type` tinyint(4) NOT NULL COMMENT '群组类型，1：小群；2：大群',
  `extra` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '附加属性',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_group_id` (`app_id`,`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='群组';

-- ----------------------------
-- Records of group
-- ----------------------------
BEGIN;
INSERT INTO `group` VALUES (1, 1, 54153362676908032, '伐木经验分享群', 'https://img2020.cnblogs.com/blog/1323675/202006/1323675-20200605171836495-1347921678.png', '', 7, 1, '', '2020-05-29 10:26:07', '2020-06-05 19:25:04');
INSERT INTO `group` VALUES (2, 1, 54153362676908033, '保卫森林交流群', 'https://img2020.cnblogs.com/blog/1323675/202006/1323675-20200605171843724-1814943720.png', ' ', 4, 1, ' ', '2020-05-29 10:26:07', '2020-06-05 17:34:05');
COMMIT;

-- ----------------------------
-- Table structure for group_user
-- ----------------------------
DROP TABLE IF EXISTS `group_user`;
CREATE TABLE `group_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) NOT NULL COMMENT 'app_id',
  `group_id` bigint(20) unsigned NOT NULL COMMENT '组id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `label` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '用户在群组的昵称',
  `extra` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '附加属性',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_group_id_user_id` (`app_id`,`group_id`,`user_id`) USING BTREE,
  KEY `idx_app_id_user_id` (`app_id`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='群组成员关系';

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` int(11) NOT NULL COMMENT 'app_id',
  `object_type` tinyint(4) NOT NULL COMMENT '所属类型，1：用户；2：群组',
  `object_id` bigint(20) unsigned NOT NULL COMMENT '所属类型的id',
  `request_id` bigint(20) NOT NULL COMMENT '请求id',
  `sender_type` tinyint(4) NOT NULL COMMENT '发送者类型',
  `sender_id` bigint(20) unsigned NOT NULL COMMENT '发送者id',
  `sender_device_id` bigint(20) unsigned NOT NULL COMMENT '发送设备id',
  `receiver_type` tinyint(4) NOT NULL COMMENT '接收者类型,1:个人；2：普通群组；3：超大群组',
  `receiver_id` bigint(20) unsigned NOT NULL COMMENT '接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id',
  `to_user_ids` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '需要@的用户id列表，多个用户用，隔开',
  `type` tinyint(4) NOT NULL COMMENT '消息类型',
  `content` varchar(4094) COLLATE utf8mb4_bin NOT NULL COMMENT '消息内容',
  `seq` bigint(20) unsigned NOT NULL COMMENT '消息序列号',
  `send_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '消息发送时间',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '消息状态，0：未处理1：消息撤回',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_object_seq` (`app_id`,`object_type`,`object_id`,`seq`) USING BTREE,
  KEY `idx_request_id` (`request_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='消息';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint(20) unsigned NOT NULL COMMENT 'app_id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户id',
  `nickname` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `sex` tinyint(4) NOT NULL COMMENT '性别，0:未知；1:男；2:女',
  `avatar_url` varchar(200) COLLATE utf8mb4_bin NOT NULL COMMENT '用户头像链接',
  `extra` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT '附加属性',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_app_id_user_id` (`app_id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1, 54146910402904064, '熊大', 1, 'https://img2020.cnblogs.com/blog/1323675/202006/1323675-20200605171809292-1360535226.jpg', '做熊，就要有个熊样', '2020-05-29 10:00:29', '2020-06-05 18:21:56');
INSERT INTO `user` VALUES (2, 1, 54146910402904065, '熊二', 1, 'https://img2020.cnblogs.com/blog/1323675/202006/1323675-20200605171821524-1987592313.jpg', '俺要吃蜂蜜', '2020-05-29 10:00:29', '2020-06-05 17:22:43');
INSERT INTO `user` VALUES (3, 1, 54146910402904066, '光头强', 1, 'https://img2020.cnblogs.com/blog/1323675/202006/1323675-20200605171828104-535094423.jpg', '臭狗熊，我饶不了你们！', '2020-05-29 10:00:29', '2020-06-05 17:22:43');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
