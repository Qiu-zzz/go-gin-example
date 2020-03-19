/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 19/03/2020 08:35:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_comment
-- ----------------------------
DROP TABLE IF EXISTS `blog_comment`;
CREATE TABLE `blog_comment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8 COLLATE utf8_bin NULL COMMENT '内容',
  `created_on` int(10) NULL DEFAULT NULL COMMENT '创建时间',
  `created_by` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '创建者名字',
  `modified_on` int(10) NULL DEFAULT NULL COMMENT '修改时间',
  `parent_id` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '父级评论id',
  `deleted_on` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
  `article_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '文章id',
  `like_count` int(10) NULL DEFAULT 0 COMMENT '点赞数',
  `dislike_count` int(10) NULL DEFAULT 0 COMMENT '点踩数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`created_by`) USING BTREE,
  INDEX `ArticleId`(`article_id`) USING BTREE,
  CONSTRAINT `ArticleId` FOREIGN KEY (`article_id`) REFERENCES `blog_article` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
