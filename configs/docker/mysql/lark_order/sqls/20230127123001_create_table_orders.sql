-- https://opendocs.alipay.com/open/204/105302?ref=api

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `order_id` bigint unsigned NOT NULL COMMENT 'order_id',
  `order_no` char(32) NOT NULL COMMENT '订单号',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
  `subject` varchar(256) NOT NULL COMMENT '订单标题',
  `time_expire` bigint unsigned NOT NULL DEFAULT '0' COMMENT '绝对超时时间',
  `total_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `pay_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '应付总额',
  `coupon_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '使用的优惠券',
  `coupon_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '优惠券抵扣金额',
  `promotion_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '促销优惠金额（促销价、满减、阶梯价）',
  `integration_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '积分抵扣金额',
  `use_integration` bigint unsigned NOT NULL DEFAULT '0' COMMENT '下单时使用的积分',
  `discount_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '后台调整订单使用的折扣金额',
  `pay_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式【1->支付宝；2->微信；3->银联；】',
  `source_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '订单来源【0->PC订单；1->app订单】',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态【0->待付款；1->已完成；2->已关闭；3->已经失效】',
  `integration` bigint unsigned NOT NULL DEFAULT '0' COMMENT '可以获得的积分',
  `growth` bigint unsigned NOT NULL DEFAULT '0' COMMENT '可以获得的成长值',
  `payment_ts` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';