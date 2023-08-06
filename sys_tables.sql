/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : test_demo

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 06/08/2023 01:44:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_tables
-- ----------------------------
DROP TABLE IF EXISTS `sys_tables`;
CREATE TABLE `sys_tables` (
  `table_id` bigint NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) DEFAULT NULL,
  `table_comment` varchar(255) DEFAULT NULL,
  `class_name` varchar(255) DEFAULT NULL,
  `tpl_category` varchar(255) DEFAULT NULL,
  `package_name` varchar(255) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL,
  `module_front_name` varchar(255) DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) DEFAULT NULL,
  `function_name` varchar(255) DEFAULT NULL,
  `function_author` varchar(255) DEFAULT NULL,
  `pk_column` varchar(255) DEFAULT NULL,
  `pk_go_field` varchar(255) DEFAULT NULL,
  `pk_json_field` varchar(255) DEFAULT NULL,
  `options` varchar(255) DEFAULT NULL,
  `tree_code` varchar(255) DEFAULT NULL,
  `tree_parent_code` varchar(255) DEFAULT NULL,
  `tree_name` varchar(255) DEFAULT NULL,
  `tree` tinyint(1) DEFAULT '0',
  `crud` tinyint(1) DEFAULT '1',
  `remark` varchar(255) DEFAULT NULL,
  `is_data_scope` tinyint DEFAULT NULL,
  `is_actions` tinyint DEFAULT NULL,
  `is_auth` tinyint DEFAULT NULL,
  `is_logical_delete` varchar(1) DEFAULT NULL,
  `logical_delete` tinyint(1) DEFAULT NULL,
  `logical_delete_column` varchar(128) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`),
  KEY `idx_sys_tables_deleted_at` (`deleted_at`),
  KEY `idx_sys_tables_create_by` (`create_by`),
  KEY `idx_sys_tables_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
