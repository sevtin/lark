-- SPU：标准化产品单元
-- SKU：库存量单位
DROP TABLE IF EXISTS `order_goods`;
CREATE TABLE `order_goods` (
  `order_goods_id` bigint unsigned NOT NULL COMMENT 'order_goods_id',
  `order_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'order_id',
  `order_no` char(32) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'order_no',
  `spu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'spu_id',
  `spu_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'spu_name',
  `spu_pic` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'spu_pic',
  `category_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品分类id',
  `sku_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品sku编号',
  `sku_name` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品sku名字',
  `sku_pic` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品sku图片',
  `sku_price` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品sku价格',
  `sku_quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '商品购买的数量',
  `sku_attrs_vals` varchar(500) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品销售属性组合（JSON）',
  `promotion_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品促销分解金额',
  `coupon_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '优惠券优惠分解金额',
  `integration_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '积分优惠分解金额',
  `real_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '该商品经过优惠后的分解金额',
  `gift_integration` int unsigned NOT NULL DEFAULT '0' COMMENT '赠送积分',
  `gift_growth` int unsigned NOT NULL DEFAULT '0' COMMENT '赠送成长值',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单项信息';
