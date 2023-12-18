DROP TABLE IF EXISTS `wallets`;
CREATE TABLE `wallets` (
  `wallet_id` bigint unsigned NOT NULL COMMENT '钱包唯一ID',
  `wallet_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '钱包类型 1-货币 单位(分) 2-钻石 3-金币 4-银币 5-铜币 6-积分',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `balance` bigint NOT NULL DEFAULT '0' COMMENT '可用余额(balance+frozen_amount=总额)(分)',
  `frozen_amount` bigint NOT NULL DEFAULT '0' COMMENT '冻结金额(分)',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '钱包状态',
  `pay_password` varchar(32) NOT NULL DEFAULT '' COMMENT '支付密码',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`wallet_id`),
  UNIQUE KEY `uid_walletType_deletedTs` (`uid`,`wallet_type`,`deleted_ts`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='钱包表';

DROP TABLE IF EXISTS `red_envelopes`;
CREATE TABLE `red_envelopes` (
  `env_id` bigint unsigned NOT NULL COMMENT '红包ID',
  `env_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '红包类型 1-均分红包 2-碰运气红包',
  `wallet_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '红包支出钱包ID',
  `receiver_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '接收者类型 1-私聊对方 2-群聊所有人 3-群聊指定人',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '交易编号',
  `chat_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'ChatID',
  `sender_uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '发红包用户ID',
  `sender_platform` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '发送平台',
  `total` bigint unsigned NOT NULL DEFAULT '0' COMMENT '红包总金额(分)',
  `quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '红包数量',
  `remain_quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '剩余红包数量',
  `remain_amount` int unsigned NOT NULL DEFAULT '0' COMMENT '剩余红包金额(分)',
  `message` varchar(128) NOT NULL DEFAULT '恭喜发财' COMMENT '祝福语',
  `env_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0-创建 1-已发放 2-已领完 3-已过期且退还剩余红包',
  `pay_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态 0-未支付 1-支付中 2-已支付 3-支付失败',
  `expired_ts` bigint unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  `finished_ts` bigint NOT NULL DEFAULT '0' COMMENT '红包领完时间',
  `receivers` varchar(1024) NOT NULL DEFAULT '' COMMENT '接收人IDs 逗号分隔',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`env_id`),
  UNIQUE KEY `uniq_tradeNo` (`trade_no`),
  KEY `idx_senderUid` (`sender_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='红包表';

DROP TABLE IF EXISTS `red_envelope_receivers`;
CREATE TABLE `red_envelope_receivers` (
  `receiver_id` bigint unsigned NOT NULL COMMENT '领取ID',
  `env_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '红包ID',
  `receiver_uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '领取用户ID',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`receiver_id`),
  KEY `idx_envId` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='红包指定接收人表';

DROP TABLE IF EXISTS `red_envelope_records`;
CREATE TABLE `red_envelope_records` (
  `record_id` bigint unsigned NOT NULL COMMENT '红包领取记录ID',
  `receiver_uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '领取用户ID',
  `env_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '红包ID',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '交易编号',
  `receive_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '领取金额(分)',
  `remain_amount` int unsigned NOT NULL DEFAULT '0' COMMENT '红包剩余金额(分)',
  `remain_quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '红包剩余数量',
  `receive_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '领取状态 0-领取中 1-成功领取 2-领取失败',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`record_id`),
  UNIQUE KEY `uniq_tradeNo` (`trade_no`),
  KEY `idx_receiverUid_envId` (`receiver_uid`,`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='红包领取记录表';

DROP TABLE IF EXISTS `fund_flows`;
CREATE TABLE `fund_flows` (
  `flow_id` bigint unsigned NOT NULL COMMENT '流水ID',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `wallet_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收/支钱包ID',
  `wallet_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '钱包类型 1-货币 单位(分) 2-钻石 3-金币 4-银币 5-铜币 6-积分',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '自编唯一交易编号',
  `assoc_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '关联ID',
  `trade_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '收支类型 1-收入 2-支出',
  `trade_type_id` int unsigned NOT NULL DEFAULT '0' COMMENT '交易类型ID',
  `trade_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '交易金额',
  `balance` bigint unsigned NOT NULL DEFAULT '0' COMMENT '交易前账户余额 暂时不用',
  `pay_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态 0-未支付 1-支付中 2-已支付 3-支付失败',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '描述信息',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`flow_id`),
  KEY `idx_tradeNo` (`trade_no`),
  KEY `idx_uid_tradeTypeId_assocId` (`uid`,`trade_type_id`,`assoc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资金流水表';

DROP TABLE IF EXISTS `trade_types`;
CREATE TABLE `trade_types` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'trade type id', -- 1-发放红包 2-领取红包 3-红包回收
  `trade_name` varchar(128) NOT NULL DEFAULT '' COMMENT '交易名称',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '描述信息',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='交易类型';

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `order_id` bigint unsigned NOT NULL COMMENT 'order id',
  `order_sn` varchar(64) NOT NULL DEFAULT '' COMMENT '订单号',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '自编唯一交易编号',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
  `time_expire` bigint unsigned NOT NULL DEFAULT '0' COMMENT '绝对超时时间',
  `amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `pay_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式 1-支付宝 2-微信 3-银联 4-PayPal',
  `source_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '订单来源 1-IOS 2-ANDROID 3-MAC 4-WINDOWS 5-WEB',
  `order_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态 0-PENDING 1-PAID 2-CANCELLED 3-REFUNDED',
  `integration` bigint unsigned NOT NULL DEFAULT '0' COMMENT '可以获得的积分',
  `growth` bigint unsigned NOT NULL DEFAULT '0' COMMENT '可以获得的成长值',
  `payment_ts` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `subject` varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '订单标题',
  `note` varchar(500) NOT NULL DEFAULT '' COMMENT '订单备注',
  `tag_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Tag ID 用于取消',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_id`),
  UNIQUE KEY `uniq_orderSn` (`order_sn`),
  UNIQUE KEY `uniq_tradeNo` (`trade_no`),
  UNIQUE KEY `uniq_tagId` (`tag_id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单';

DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
  `order_item_id` bigint unsigned NOT NULL COMMENT 'order item id',
  `order_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'order id',
  `spu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'spu id',
  `spu_name` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'spu name',
  `spu_pic` varchar(160) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品spu图片',
  `cat_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品分类id',
  `sku_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'sku id',
  `sku_name` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'sku name',
  `sku_pic` varchar(160) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品sku图片',
  `sku_price` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品sku价格',
  `sku_quantity` int unsigned NOT NULL DEFAULT '0' COMMENT '商品购买的数量',
  `sku_attrs` varchar(1000) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '商品销售属性组合（JSON）',
  `promotion_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品促销分解金额',
  `coupon_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '优惠券优惠分解金额',
  `integration_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '积分优惠分解金额',
  `real_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '该商品经过优惠后的分解金额',
  `gift_integration` int unsigned NOT NULL DEFAULT '0' COMMENT '赠送积分',
  `gift_growth` int unsigned NOT NULL DEFAULT '0' COMMENT '赠送成长值',
  `created_ts` bigint unsigned NOT NULL DEFAULT '0',
  `updated_ts` bigint unsigned NOT NULL DEFAULT '0',
  `deleted_ts` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_item_id`),
  KEY `idx_orderId` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单项信息';

DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
  `pay_id` bigint unsigned NOT NULL COMMENT 'pay id',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
  `seller_id` varchar(64) NOT NULL DEFAULT '' COMMENT '收款账号对应的第三方支付唯一用户号',
  `buyer_id` varchar(64) NOT NULL DEFAULT '' COMMENT '支付人所在支付平台ID',
  `buyer_email` varchar(64) NOT NULL DEFAULT '' COMMENT '支付人所在支付平台Email',
  `order_id` bigint NOT NULL DEFAULT '0' COMMENT '商户原始订单号',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '系统内部唯一订单号，只能是数字、大小写字母_-*',
  `pay_status` tinyint NOT NULL DEFAULT '0' COMMENT '支付状态 0-待支付 1-已支付/已完成 2-已取消 3-失败',
  `currency` varchar(10) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '币种',
  `subject` varchar(256) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '订单标题',
  `summary` varchar(256) NOT NULL DEFAULT '' COMMENT '摘要',
  `th_trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '该交易在第三方支付系统中的交易流水号',
  `pay_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付方式 1-支付宝 2-微信 3-银联 4-Paypal',
  `return_content` varchar(4096) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'return内容',
  `notify_content` varchar(4096) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'notify内容',
  `result_content` varchar(4096) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '结果内容',
  `return_ts` bigint NOT NULL DEFAULT '0' COMMENT 'return时间',
  `notify_ts` bigint NOT NULL DEFAULT '0' COMMENT 'notify时间',
  `total_amount` bigint NOT NULL DEFAULT '0' COMMENT '订单金额',
  `actual_amount` bigint NOT NULL DEFAULT '0' COMMENT '实际入账金额',
  `pay_ts` bigint NOT NULL DEFAULT '0' COMMENT '支付时间',
  `tag_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Tag ID 用于取消',
  `sale_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Sale ID 用于退款',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`pay_id`),
  UNIQUE KEY `uniq_orderId` (`order_id`),
  KEY `idx_tagId` (`tag_id`),
  KEY `idx_saleId` (`sale_id`),
  KEY `idx_tradeNo` (`trade_no`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支付';