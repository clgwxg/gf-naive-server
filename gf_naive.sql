/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : gf_admin

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 03/04/2024 10:30:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门名称',
  `parent_id` int NOT NULL COMMENT '上级部门id',
  `leader` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门负责人',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` tinyint NULL DEFAULT 1 COMMENT '部门状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dept
-- ----------------------------
INSERT INTO `dept` VALUES (1, 'gf-naive', 0, NULL, NULL, NULL, 1, '2024-03-23 19:31:30', '2024-03-23 19:31:38');
INSERT INTO `dept` VALUES (7, '郑州分公司', 1, '王五', '', '', 1, '2024-03-23 11:35:15', '2024-03-23 11:35:15');
INSERT INTO `dept` VALUES (8, '研发部门', 7, '', '', '', 1, '2024-03-23 11:43:32', '2024-03-23 11:43:32');

-- ----------------------------
-- Table structure for emp
-- ----------------------------
DROP TABLE IF EXISTS `emp`;
CREATE TABLE `emp`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `nick_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户昵称',
  `dept_id` int NULL DEFAULT NULL COMMENT '部门id',
  `role_id` int NULL DEFAULT NULL COMMENT '用户角色',
  `post_id` int NULL DEFAULT NULL COMMENT '用户岗位',
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `email` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint NULL DEFAULT NULL COMMENT '用户状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of emp
-- ----------------------------
INSERT INTO `emp` VALUES (3, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '张三', 8, 34, 6, '13000000000', '', '', 1, '2024-03-23 11:44:54', '2024-03-23 11:58:32');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` int NOT NULL DEFAULT 0 COMMENT '父级id',
  `type` tinyint NULL DEFAULT NULL COMMENT '菜单类型：1目录2菜单3按钮',
  `icon` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `url` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由地址',
  `is_frame` tinyint NULL DEFAULT NULL COMMENT '是否外链：1是0否',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由参数',
  `is_cache` tinyint NULL DEFAULT NULL COMMENT '是否缓存：1缓存0不缓存',
  `perms` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api权限',
  `visible` tinyint NULL DEFAULT NULL COMMENT '是否显示：1显示0不显示',
  `status` tinyint NULL DEFAULT NULL COMMENT '菜单状态：1正常0禁用',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, '系统管理', 0, 1, 'system', '', 0, '', '', 0, '', 1, 1, '2024-03-23 10:49:32', '2024-03-23 10:49:32', NULL);
INSERT INTO `menu` VALUES (2, '菜单管理', 1, 2, 'menu', '/system/menu', 0, '/system/MenuView/index', '', 0, 'system:menu:list', 1, 1, '2024-03-23 10:53:44', '2024-03-29 14:14:35', NULL);
INSERT INTO `menu` VALUES (3, '新增', 2, 3, '', '', 0, '', '', 0, 'system:menu:add', 1, 1, '2024-03-23 10:55:57', '2024-03-29 14:14:50', NULL);
INSERT INTO `menu` VALUES (4, '编辑', 2, 3, '', '', 0, '', '', 0, 'system:menu:edit', 1, 1, '2024-03-23 10:56:37', '2024-03-29 14:15:06', NULL);
INSERT INTO `menu` VALUES (5, '删除', 2, 3, '', '', 0, '', '', 0, 'system:menu:delete', 1, 1, '2024-03-23 10:57:18', '2024-03-29 14:15:16', NULL);
INSERT INTO `menu` VALUES (6, '查看', 2, 3, '', '', 0, '', '', 0, 'system:menu:view', 1, 1, '2024-03-23 10:58:06', '2024-03-29 14:15:29', NULL);
INSERT INTO `menu` VALUES (7, '角色管理', 1, 2, 'role', '/system/role', 0, '/system/RoleView/index', '', 0, 'system:role:list', 1, 1, '2024-03-23 10:59:41', '2024-03-29 14:23:27', NULL);
INSERT INTO `menu` VALUES (19, '新增', 7, 3, '', '', 0, '', '', 0, 'system:role:add', 1, 1, '2024-03-23 11:04:58', '2024-03-29 14:23:45', NULL);
INSERT INTO `menu` VALUES (20, '编辑', 7, 3, '', '', 0, '', '', 0, 'system:role:edit', 1, 1, '2024-03-23 11:05:14', '2024-03-29 14:23:57', NULL);
INSERT INTO `menu` VALUES (21, '删除', 7, 3, '', '', 0, '', '', 0, 'system:role:delete', 1, 1, '2024-03-23 11:05:38', '2024-03-29 14:24:10', NULL);
INSERT INTO `menu` VALUES (22, '查看', 7, 3, '', '', 0, '', '', 0, 'system:role:view', 1, 1, '2024-03-23 11:05:52', '2024-03-29 14:24:21', NULL);
INSERT INTO `menu` VALUES (23, '员工管理', 1, 2, 'emp', '/system/emp', 0, '/system/EmpView/index', '', 0, 'system:emp:list', 1, 1, '2024-03-23 11:08:52', '2024-03-29 14:24:49', NULL);
INSERT INTO `menu` VALUES (24, '新增', 23, 3, '', '', 0, '', '', 0, 'system:emp:add', 1, 1, '2024-03-23 11:09:10', '2024-03-29 14:25:03', NULL);
INSERT INTO `menu` VALUES (25, '编辑', 23, 3, '', '', 0, '', '', 0, 'system:emp:edit', 1, 1, '2024-03-23 11:09:27', '2024-03-29 14:25:14', NULL);
INSERT INTO `menu` VALUES (26, '删除', 23, 3, '', '', 0, '', '', 0, 'system:emp:delete', 1, 1, '2024-03-23 11:09:50', '2024-03-29 14:25:28', NULL);
INSERT INTO `menu` VALUES (27, '查看', 23, 3, '', '', 0, '', '', 0, 'system:emp:view', 1, 1, '2024-03-23 11:10:02', '2024-03-29 14:25:38', NULL);
INSERT INTO `menu` VALUES (28, '部门管理', 1, 2, 'dept', '/system/dept', 0, '/system/DeptView/index', '', 0, 'system:dept:list', 1, 1, '2024-03-23 11:12:23', '2024-03-29 14:25:58', NULL);
INSERT INTO `menu` VALUES (29, '新增', 28, 3, '', '', 0, '', '', 0, 'system:dept:add', 1, 1, '2024-03-23 11:12:33', '2024-03-29 14:26:14', NULL);
INSERT INTO `menu` VALUES (30, '编辑', 28, 3, '', '', 0, '', '', 0, 'system:dept:edit', 1, 1, '2024-03-23 11:12:44', '2024-03-29 14:26:25', NULL);
INSERT INTO `menu` VALUES (31, '删除', 28, 3, '', '', 0, '', '', 0, 'system:dept:delete', 1, 1, '2024-03-23 11:13:02', '2024-03-29 14:26:35', NULL);
INSERT INTO `menu` VALUES (32, '查看', 28, 3, '', '', 0, '', '', 0, 'system:dept:view', 1, 1, '2024-03-23 11:13:12', '2024-03-29 14:26:44', NULL);
INSERT INTO `menu` VALUES (33, '岗位管理', 1, 2, 'post', '/system/post', 0, '/system/PostView/index', '', 0, 'system:post:list', 1, 1, '2024-03-23 11:15:56', '2024-03-29 14:27:00', NULL);
INSERT INTO `menu` VALUES (34, '新增', 33, 3, '', '', 0, '', '', 0, 'system:post:add', 1, 1, '2024-03-23 11:16:09', '2024-03-29 14:27:37', NULL);
INSERT INTO `menu` VALUES (35, '编辑', 33, 3, '', '', 0, '', '', 0, 'system:post:edit', 1, 1, '2024-03-23 11:16:21', '2024-03-29 14:27:45', NULL);
INSERT INTO `menu` VALUES (36, '删除', 33, 3, '', '', 0, '', '', 0, 'system:post:delete', 1, 1, '2024-03-23 11:16:36', '2024-03-29 14:27:55', NULL);
INSERT INTO `menu` VALUES (37, '查看', 33, 3, '', '', 0, '', '', 0, 'system:post:view', 1, 1, '2024-03-23 11:16:51', '2024-03-29 14:28:05', NULL);

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位名称',
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '岗位编码',
  `status` int NULL DEFAULT 1 COMMENT '岗位状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES (5, '董事长', 'ceo', 1, '', '2024-03-23 11:38:25', '2024-03-23 11:38:25');
INSERT INTO `post` VALUES (6, '项目经理', 'se', 1, '', '2024-03-23 11:40:00', '2024-03-23 11:40:00');
INSERT INTO `post` VALUES (7, '人力资源', 'hr', 1, '', '2024-03-23 11:40:21', '2024-03-23 11:40:21');
INSERT INTO `post` VALUES (8, '普通员工', 'user', 1, '', '2024-03-23 11:40:39', '2024-03-23 11:40:39');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `admin` tinyint NOT NULL DEFAULT 0 COMMENT '是否超级管理员：1是0否',
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注',
  `status` tinyint NULL DEFAULT 0 COMMENT '角色状态：1正常0停用',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, 34, '', 1, '超级管理员', '2024-03-23 11:19:01', '2024-03-23 11:19:01');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_id` int NOT NULL,
  `menu_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (35, 34, 1);
INSERT INTO `role_menu` VALUES (36, 34, 2);
INSERT INTO `role_menu` VALUES (37, 34, 3);
INSERT INTO `role_menu` VALUES (38, 34, 7);
INSERT INTO `role_menu` VALUES (39, 34, 23);
INSERT INTO `role_menu` VALUES (40, 34, 28);
INSERT INTO `role_menu` VALUES (41, 34, 33);
INSERT INTO `role_menu` VALUES (42, 34, 4);
INSERT INTO `role_menu` VALUES (43, 34, 5);
INSERT INTO `role_menu` VALUES (44, 34, 6);
INSERT INTO `role_menu` VALUES (45, 34, 19);
INSERT INTO `role_menu` VALUES (46, 34, 20);
INSERT INTO `role_menu` VALUES (47, 34, 21);
INSERT INTO `role_menu` VALUES (48, 34, 22);
INSERT INTO `role_menu` VALUES (49, 34, 24);
INSERT INTO `role_menu` VALUES (50, 34, 25);
INSERT INTO `role_menu` VALUES (51, 34, 26);
INSERT INTO `role_menu` VALUES (52, 34, 27);
INSERT INTO `role_menu` VALUES (53, 34, 29);
INSERT INTO `role_menu` VALUES (54, 34, 30);
INSERT INTO `role_menu` VALUES (55, 34, 31);
INSERT INTO `role_menu` VALUES (56, 34, 32);
INSERT INTO `role_menu` VALUES (57, 34, 34);
INSERT INTO `role_menu` VALUES (58, 34, 35);
INSERT INTO `role_menu` VALUES (59, 34, 36);
INSERT INTO `role_menu` VALUES (60, 34, 37);

SET FOREIGN_KEY_CHECKS = 1;
