DROP TABLE IF EXISTS `wallets`;
CREATE TABLE `wallets` (
  `wallet_id` bigint unsigned NOT NULL COMMENT '钱包唯一ID',
  `wallet_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '钱包类型',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `balance` bigint NOT NULL DEFAULT '0' COMMENT '可用余额(balance+frozen_amount=总额)(分)',
  `frozen_amount` bigint NOT NULL DEFAULT '0' COMMENT '冻结金额(分)',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '钱包状态',
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
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '自编唯一交易编号',
  `assoc_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '关联ID',
  `trade_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '收支类型 1-收入 2-支出',
  `trade_type_id` int unsigned NOT NULL DEFAULT '0' COMMENT '交易类型ID',
  `trade_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '交易金额',
  `balance` bigint unsigned NOT NULL DEFAULT '0' COMMENT '交易前账户余额',
  `pay_status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态 0-未支付 1-支付中 2-已支付 3-支付失败',
  `description` varchar(500) NOT NULL DEFAULT '' COMMENT '描述信息',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`flow_id`),
  UNIQUE KEY `uniq_tradeNo` (`trade_no`),
  KEY `idx_uid_tradeTypeId_assocId` (`uid`,`trade_type_id`,`assoc_id`),
  KEY `idx_uid_tradeNo` (`uid`,`trade_no`)
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
