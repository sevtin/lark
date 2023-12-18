-- 第三方支付
DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
  `payment_id` bigint unsigned NOT NULL COMMENT 'payment_id',
  `seller_id` varchar(32) NOT NULL DEFAULT '' COMMENT '收款账号对应的第三方支付唯一用户号',
  `order_id` bigint NOT NULL DEFAULT '0' COMMENT '商户原始订单号',
  `order_sn` varchar(32) NOT NULL DEFAULT '' COMMENT '系统内部唯一订单号，只能是数字、大小写字母_-*',
  `payment_status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态 0:待支付 1:已支付/已完成 2:已取消 3:失败',
  `subject` varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '订单标题',
  `trade_no` varchar(32) NOT NULL DEFAULT '' COMMENT '该交易在第三方支付系统中的交易流水号',
  `callback_content` varchar(4096) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '回调内容',
  `callback_ts` bigint NOT NULL DEFAULT '0' COMMENT '回调时间',
  `total_amount` bigint NOT NULL DEFAULT '0' COMMENT '订单金额',
  `actual_amount` bigint NOT NULL DEFAULT '0' COMMENT '实际入账金额',
  `confirm_ts` bigint NOT NULL DEFAULT '0' COMMENT '确认时间',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`payment_id`),
  UNIQUE KEY `uniq_order_id` (`order_id`),
  UNIQUE KEY `uniq_order_sn` (`order_sn`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='支付';
