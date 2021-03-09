CREATE DATABASE IF NOT EXISTS `myblog` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `myblog`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog
-- ----------------------------
DROP TABLE IF EXISTS `blog`;
CREATE TABLE `blog`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `summary` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '摘要',
  `contend` longblob NULL COMMENT '内容',
  `category_id` int(11) UNSIGNED NOT NULL COMMENT '分类id',
  `tag_ids` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签ids',
  `hits` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击量',
  `image_url` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片url',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `in_of_create`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 610 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `count` int(11) NOT NULL COMMENT '数量',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

DROP TABLE IF EXISTS `db_version`;
CREATE TABLE `db_version`  (
  `db_version` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0.0.0.0'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;


DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '盐',
  `nick_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_mail` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮件地址',
  `privileges` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限',
  `login_token` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录Token',
  `login_token_time` datetime NULL DEFAULT NULL COMMENT '登录Token生成时间',
  `state` smallint(2) UNSIGNED NOT NULL COMMENT '状态（0：初始，1：禁用，2：正常）',
  `parent_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '上级',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '修改时间',
  `delete_time` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

INSERT INTO `user` VALUES (1, 'admin', '0180cede84ad298b371f5506375ce52d', '198793', '超级管理员', 'liu865488158@qq.com', 'myProject,gmmx,sbtj,sbtj_unity,gmmx_win,zjzr,zjzr2,zjzr_cocos,zjzr2_unity,gmmx_unity,zjzr2_win,system,refreshdbcache,systeminfo,buildFlow,flowmanager,buildrecord,index', '', NULL, 2, '', '2021-01-28 11:17:33', '2021-01-28 11:17:37', NULL);
INSERT INTO `db_version` VALUES ('0.0.0.0');

INSERT INTO privilege(label,path,`name`,icon,is_leaf,is_forbidden,created_at,updated_at) VALUES('菜单管理','/menuManage','menuManage','icon',0,0,now(),now());
INSERT INTO privilege(label,path,`name`,icon,is_leaf,is_forbidden,created_at,updated_at) VALUES('用户管理','/userManage','userManage','icon',0,0,now(),now());
INSERT INTO privilege(label,path,`name`,icon,is_leaf,is_forbidden,created_at,updated_at) VALUES('测试管理','/testManage','testManage','icon',0,0,now(),now());
INSERT INTO privilege(label,path,`name`,icon,is_leaf,is_forbidden,parent_name,parent_path,created_at,updated_at) VALUES('测试子管理','/testManage1','testManage1','icon',1,0,'testManage','/testManage',now(),now());
INSERT INTO `db_version` VALUES ('0.0.0.1');
