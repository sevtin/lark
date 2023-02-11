DROP TABLE IF EXISTS `wallets`;
CREATE TABLE `wallets` (
  `wallet_id` bigint NOT NULL DEFAULT '0' COMMENT '钱包唯一ID',
  `wallet_type` tinyint NOT NULL DEFAULT '0' COMMENT '钱包类型',
  `uid` bigint NOT NULL DEFAULT '0' COMMENT '用户UID',
  `balance` bigint NOT NULL DEFAULT '0' COMMENT '可用余额(balance+frozen_amount=总额)',
  `frozen_amount` bigint NOT NULL DEFAULT '0' COMMENT '冻结金额',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '钱包状态',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`wallet_id`),
  UNIQUE KEY `walletType_uid_deletedTs` (`wallet_type`,`uid`,`deleted_ts`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_walletType` (`wallet_type`),
  KEY `idx_uid` (`uid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;