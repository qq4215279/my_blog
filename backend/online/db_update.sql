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
  `contend` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '内容',
  `category_id` int(11) UNSIGNED NOT NULL COMMENT '分类id',
  `tag_ids` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签ids',
  `hits` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击量',
  `image_url` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '图片url',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `in_of_create`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of blog
-- ----------------------------

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  `count` int(11) NOT NULL COMMENT '数量',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of category
-- ----------------------------

-- ----------------------------
-- Table structure for db_version
-- ----------------------------
DROP TABLE IF EXISTS `db_version`;
CREATE TABLE `db_version`  (
  `db_version` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0.0.0.0'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of db_version
-- ----------------------------
INSERT INTO `db_version` VALUES ('0.0.0.1');

-- ----------------------------
-- Table structure for privilege
-- ----------------------------
DROP TABLE IF EXISTS `privilege`;
CREATE TABLE `privilege`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `label` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '展示名称',
  `path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'path',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '模块名称',
  `icon` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '图标',
  `is_leaf` tinyint(3) NOT NULL COMMENT '是否是叶子节点（0：不是，1：是）',
  `is_forbidden` tinyint(3) NOT NULL DEFAULT 0 COMMENT '是否禁用(0:未禁用，1：禁用)',
  `parent_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '父级节点名',
  `parent_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '父级节路径',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of privilege
-- ----------------------------
INSERT INTO `privilege` VALUES (1, '菜单管理', '/menuManage', 'menuManage', 'icon', 0, 0, '', '', now(), now(), NULL);
INSERT INTO `privilege` VALUES (2, '用户管理', '/userManage', 'userManage', 'icon', 0, 0, '', '', now(), now(), NULL);
INSERT INTO `privilege` VALUES (3, '测试管理', '/testManage', 'testManage', 'icon', 0, 0, '', '', now(), now(), NULL);
INSERT INTO `privilege` VALUES (4, '测试子管理', '/testManage1', 'testManage1', 'icon', 1, 0, 'testManage', '/testManage', now(), now(), NULL);
INSERT INTO `privilege` VALUES (5, '文章', '/article', 'article', 'icon', 0, 0, '', '', now(), now(), NULL);
INSERT INTO `privilege` VALUES (6, '评论管理', '/article/commentManage', 'commentManage', 'icon', 1, 0, 'article', '/article', now(), now(), NULL);
INSERT INTO `privilege` VALUES (7, '标签管理', '/article/tagManage', 'tagManage', 'icon', 1, 0, 'article', '/article', now(), now(), NULL);
INSERT INTO `privilege` VALUES (8, '文章管理', '/article/articleManage', 'articleManage', 'icon', 1, 0, 'article', '/article', now(), now(), NULL);
INSERT INTO `privilege` VALUES (9, '通知', '/noticeManage', 'noticeManage', 'icon', 0, 0, '', '', now(), now(), NULL);
-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名称',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unq_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tag
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
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
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '0180cede84ad298b371f5506375ce52d', '198793', '超级管理员', 'liu865488158@qq.com', 'myProject,gmmx,sbtj,sbtj_unity,gmmx_win,zjzr,zjzr2,zjzr_cocos,zjzr2_unity,gmmx_unity,zjzr2_win,system,refreshdbcache,systeminfo,buildFlow,flowmanager,buildrecord,index', '3GLTIpSqsGU0uFhrGBAl', '0000-00-00 00:00:00', 2, '', '2021-01-28 11:17:33', '2021-03-23 21:17:29', NULL);
SET FOREIGN_KEY_CHECKS = 1;

UPDATE db_version SET db_version = '0.0.0.2';
