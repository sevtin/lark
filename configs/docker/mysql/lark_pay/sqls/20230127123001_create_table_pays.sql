-- https://opendocs.alipay.com/open/204/105302?ref=api

DROP TABLE IF EXISTS `pays`;
CREATE TABLE `pays` (
  `pay_id` bigint unsigned NOT NULL COMMENT 'pay_id',
  `app_id` char(32) NOT NULL COMMENT '配给开发者的应用ID',
  `out_trade_no` char(64) NOT NULL COMMENT '商户网站唯一订单号',
  `merchant_order_no` varchar(32) NOT NULL COMMENT '商户原始订单号',
  `total_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `subject` varchar(256) NOT NULL COMMENT '订单标题',
  `time_expire` bigint unsigned NOT NULL DEFAULT '0' COMMENT '绝对超时时间',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`pay_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付';