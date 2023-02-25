DROP TABLE IF EXISTS `bank_cards`;
CREATE TABLE `bank_cards` (
  `card_id` bigint NOT NULL DEFAULT '0' COMMENT '卡ID',
  `uid` bigint NOT NULL DEFAULT '0' COMMENT '用户UID',
  `bank_id` bigint NOT NULL DEFAULT '0' COMMENT '银行ID',
  `beneficiary` varchar(64) NOT NULL DEFAULT '' COMMENT '帐户名',
  `account_number` varchar(19) NOT NULL COMMENT '银行卡卡号',
  `card_type` tinyint NOT NULL DEFAULT '0' COMMENT '0:储蓄卡 1:信用卡',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '银行卡状态',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`card_id`),
  UNIQUE KEY `uid_accountNumber_deletedTs` (`uid`,`account_number`,`deleted_ts`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_uid` (`uid`),
  KEY `idx_bankId` (`bank_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
