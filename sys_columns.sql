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

 Date: 06/08/2023 01:44:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_columns
-- ----------------------------
DROP TABLE IF EXISTS `sys_columns`;
CREATE TABLE `sys_columns` (
  `column_id` bigint NOT NULL AUTO_INCREMENT,
  `table_id` bigint DEFAULT NULL,
  `column_name` varchar(128) DEFAULT NULL,
  `column_comment` varchar(128) DEFAULT NULL,
  `column_type` varchar(128) DEFAULT NULL,
  `go_type` varchar(128) DEFAULT NULL,
  `go_field` varchar(128) DEFAULT NULL,
  `json_field` varchar(128) DEFAULT NULL,
  `is_pk` varchar(4) DEFAULT NULL,
  `is_increment` varchar(4) DEFAULT NULL,
  `is_required` varchar(4) DEFAULT NULL,
  `is_insert` varchar(4) DEFAULT NULL,
  `is_edit` varchar(4) DEFAULT NULL,
  `is_list` varchar(4) DEFAULT NULL,
  `is_query` varchar(4) DEFAULT NULL,
  `query_type` varchar(128) DEFAULT NULL,
  `html_type` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `list` varchar(1) DEFAULT NULL,
  `pk` tinyint(1) DEFAULT NULL,
  `required` tinyint(1) DEFAULT NULL,
  `super_column` tinyint(1) DEFAULT NULL,
  `usable_column` tinyint(1) DEFAULT NULL,
  `increment` tinyint(1) DEFAULT NULL,
  `insert` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `query` tinyint(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `fk_table_name` longtext,
  `fk_table_name_class` longtext,
  `fk_table_name_package` longtext,
  `fk_label_id` longtext,
  `fk_label_name` varchar(255) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`column_id`),
  KEY `idx_sys_columns_deleted_at` (`deleted_at`),
  KEY `idx_sys_columns_create_by` (`create_by`),
  KEY `idx_sys_columns_update_by` (`update_by`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
