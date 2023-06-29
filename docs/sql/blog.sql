CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) DEFAULT '' COMMENT '文章标题',
    `desc` varchar(255) DEFAULT '' COMMENT '简述',
    `content` text COMMENT '文章内容',
    `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    `status` tinyint(3) unsigned DEFAULT '1' COMMENT '状态',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '文章管理';
-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
    `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

INSERT INTO `blog_service`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'hd2yao', 'go-blog', 0, 'yhd', 0, '', 0, 0);

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '标签名称',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签管理';

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `article_id` int(11) NOT NULL COMMENT '文章 ID',
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签 ID',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';