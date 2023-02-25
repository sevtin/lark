DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `goods_id` bigint unsigned NOT NULL COMMENT 'good_id',
  `goods_no` varchar(256) NOT NULL COMMENT '统一商品编号',
  `pay_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付ID',
  `goods_name` varchar(256) NOT NULL COMMENT '商品名称',
  `goods_desc` varchar(512) NOT NULL COMMENT '商品描述',
  `show_url` varchar(512) NOT NULL COMMENT '商品的展示地址',
  `quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `price` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品单价',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`goods_id`),
  KEY `idx_payId` (`pay_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品';
