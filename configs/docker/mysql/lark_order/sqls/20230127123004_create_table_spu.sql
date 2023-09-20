DROP TABLE IF EXISTS `spu_infos`;
CREATE TABLE `spu_infos` (
  `spu_id` bigint unsigned NOT NULL COMMENT 'spu id',
  `spu_name` varchar(200) NOT NULL DEFAULT '' COMMENT '商品名称',
  `spu_description` varchar(1000) NOT NULL DEFAULT '' COMMENT '商品描述',
  `cat_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '所属分类id',
  `brand_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '品牌id',
  `publish_status` tinyint(2) DEFAULT '0' COMMENT '状态 0-隐藏 1-上架 2-下架',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`spu_id`),
  KEY `idx_catId` (`cat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='spu信息';

INSERT INTO `spu_infos` (`spu_id`, `spu_name`, `spu_description`, `cat_id`, `brand_id`, `publish_status`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (1, '红包', '红包', 3, 0, 0, 1694055993, 1694055993, 0);

DROP TABLE IF EXISTS `spu_info_descs`;
CREATE TABLE `spu_info_descs` (
  `spu_desc_id` bigint unsigned NOT NULL COMMENT 'spu desc id',
  `spu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'spu id',
  `description` varchar(2000) NOT NULL DEFAULT '' COMMENT '商品介绍',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`spu_desc_id`),
  KEY `idx_spuId` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='spu信息介绍';

INSERT INTO `spu_info_descs` (`spu_desc_id`, `spu_id`, `description`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (1, 1, '红包', 1694055993, 1694055993, 0);

DROP TABLE IF EXISTS `spu_images`;
CREATE TABLE `spu_images` (
  `img_id` bigint unsigned NOT NULL COMMENT 'id',
  `spu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'spu id',
  `img_name` varchar(200) NOT NULL DEFAULT '' COMMENT '图片名',
  `img_url` varchar(160) NOT NULL DEFAULT '' COMMENT '图片地址',
  `img_sort` int unsigned NOT NULL DEFAULT '0' COMMENT '顺序',
  `default_img` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0-不是默认图 1-是默认图',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`img_id`),
  KEY `idx_spuId` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='spu图片';

INSERT INTO `spu_images` (`img_id`, `spu_id`, `img_name`, `img_url`, `img_sort`, `default_img`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (1, 1, '红包', 'http://lark-minio.com:19000/photos/6b536cc7-5e3a-4d31-8018-1e5853f88a1c.png', 0, 1, 1694055993, 1694055993, 0);

DROP TABLE IF EXISTS `pms_categories`;
CREATE TABLE `pms_categories` (
  `cat_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父分类id',
  `cat_level` int NOT NULL DEFAULT '0' COMMENT '层级',
  `show_status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否显示 0-不显示 1显示',
  `cat_sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `icon` varchar(160) NOT NULL DEFAULT '' COMMENT '图标地址',
  `product_unit` varchar(50) NOT NULL DEFAULT '' COMMENT '计量单位',
  `product_count` int NOT NULL DEFAULT '0' COMMENT '商品数量',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`cat_id`),
  KEY `idx_parentId` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品三级分类';

INSERT INTO `pms_categories` (`cat_id`, `name`, `parent_id`, `cat_level`, `show_status`, `cat_sort`, `icon`, `product_unit`, `product_count`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (1, '红包', 0, 1, 1, 0, 'http://lark-minio.com:19000/photos/6b536cc7-5e3a-4d31-8018-1e5853f88a1c.png', '封', -1, 1694055993, 1694055993, 0);
INSERT INTO `pms_categories` (`cat_id`, `name`, `parent_id`, `cat_level`, `show_status`, `cat_sort`, `icon`, `product_unit`, `product_count`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (2, '红包', 1, 2, 1, 0, 'http://lark-minio.com:19000/photos/6b536cc7-5e3a-4d31-8018-1e5853f88a1c.png', '封', -1, 1694055993, 1694055993, 0);
INSERT INTO `pms_categories` (`cat_id`, `name`, `parent_id`, `cat_level`, `show_status`, `cat_sort`, `icon`, `product_unit`, `product_count`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (3, '红包', 2, 3, 1, 0, 'http://lark-minio.com:19000/photos/6b536cc7-5e3a-4d31-8018-1e5853f88a1c.png', '封', -1, 1694055993, 1694055993, 0);