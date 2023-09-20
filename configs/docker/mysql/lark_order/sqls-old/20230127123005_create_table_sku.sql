DROP TABLE IF EXISTS `sku_infos`;
CREATE TABLE `sku_infos` (
  `sku_id` bigint unsigned NOT NULL COMMENT 'sku id',
  `spu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'spu id',
  `sku_name` varchar(200) NOT NULL DEFAULT '' COMMENT 'sku名称',
  `sku_desc` varchar(2000) NOT NULL DEFAULT '' COMMENT 'sku介绍描述',
  `cat_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '所属分类id',
  `brand_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '品牌id',
  `sku_img` varchar(160) NOT NULL DEFAULT '' COMMENT '图片',
  `sku_title` varchar(200) NOT NULL DEFAULT '' COMMENT '标题',
  `sku_subtitle` varchar(200) NOT NULL DEFAULT '' COMMENT '副标题',
  `price` bigint NOT NULL DEFAULT '0' COMMENT '价格',
  `sale_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '销量',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`sku_id`),
  KEY `idx_spuId` (`spu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='sku信息';

INSERT INTO `sku_infos` (`sku_id`, `spu_id`, `sku_name`, `sku_desc`, `cat_id`, `brand_id`, `sku_img`, `sku_title`, `sku_subtitle`, `price`, `sale_count`, `created_ts`, `updated_ts`, `deleted_ts`)
VALUES (1, 1, '红包', '红包', 3, 0, 'http://lark-minio.com:19000/photos/6b536cc7-5e3a-4d31-8018-1e5853f88a1c.png', '红包', '红包', -1, 0, 1694055993, 1694055993, 0);

DROP TABLE IF EXISTS `sku_sale_attributes`;
CREATE TABLE `sku_sale_attributes` (
  `attr_id` bigint unsigned NOT NULL COMMENT 'attr id',
  `sku_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'sku id',
  `attr_name` varchar(200) NOT NULL DEFAULT '' COMMENT '销售属性名',
  `attr_value` varchar(200) NOT NULL DEFAULT '' COMMENT '销售属性值',
  `attr_sort` int unsigned NOT NULL DEFAULT '0' COMMENT '顺序',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`attr_id`),
  KEY `idx_skuId` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='sku销售属性';

DROP TABLE IF EXISTS `sku_images`;
CREATE TABLE `sku_images` (
  `img_id` bigint unsigned NOT NULL  COMMENT 'img id',
  `sku_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'sku id',
  `img_url` varchar(160) NOT NULL DEFAULT '' COMMENT '图片地址',
  `img_sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `default_img` tinyint(2) NOT NULL DEFAULT '0' COMMENT '0-不是默认图 1-是默认图',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`img_id`),
  KEY `idx_skuId` (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='sku图片';