-- https://opendocs.alipay.com/open/204/105302?ref=api

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `order_id` bigint unsigned NOT NULL COMMENT 'order_id',
  `order_sn` char(32) NOT NULL DEFAULT '' COMMENT '订单号',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
  `subject` varchar(256) NOT NULL DEFAULT '' COMMENT '订单标题',
  `time_expire` bigint unsigned NOT NULL DEFAULT '0' COMMENT '绝对超时时间',
  `total_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `pay_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '应付总额',
  `pay_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式 1-支付宝 2-微信 3-银联',
  `source_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '订单来源 1-IOS 2-ANDROID 3-MAC 4-WINDOWS 5-WEB',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态 0-待付款 1-已完成 2-已关闭 3-已经失效',
  `payment_ts` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `note` varchar(500) NOT NULL DEFAULT '' COMMENT '订单备注',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_id`),
  KEY `idx_orderSn` (`order_sn`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';