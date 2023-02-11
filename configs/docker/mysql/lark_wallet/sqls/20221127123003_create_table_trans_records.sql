-- 交易记录表
DROP TABLE IF EXISTS `trans_records`;
CREATE TABLE `trans_records` (
  `tsn` bigint NOT NULL DEFAULT '0' COMMENT '交易序列号 Transaction Number',
  `uid` bigint NOT NULL DEFAULT '0' COMMENT '交易者UID',
  `from_wid` bigint NOT NULL DEFAULT '0' COMMENT 'from 钱包ID',
  `to_wid` bigint NOT NULL DEFAULT '0' COMMENT 'to 钱包ID',
  `to_account` varchar(32) NOT NULL COMMENT 'to 账户',
  `account_type` tinyint NOT NULL DEFAULT '0' COMMENT 'to_account账户类型 1:lark钱包 2:银行卡 3:支付宝 4:微信',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '交易额',
  `exc_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '对换率',
  `exc_amount` bigint NOT NULL DEFAULT '0' COMMENT '对换目标额',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '交易状态 0:待支付 1:已支付/已完成 2:已取消 3:失败',
  `trader_role` tinyint NOT NULL DEFAULT '0' COMMENT '交易者的角色 0:付款人 1:收款人',
  `trans_type` bigint NOT NULL DEFAULT '0' COMMENT '交易类型 1:转账支出 2:转账收入 3:兑换支出 4:兑换收入 5:交易支付 6:交易收款 7:提现',

  `appid` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT 'APPID',
  `mchId` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '商户ID(不对前端开放)',
  `out_trade_no` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '系统内部唯一订单号，只能是数字、大小写字母_-*',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '该交易在第三方支付系统中的交易流水号',
  `description` varchar(200) CHARACTER SET utf8mb4 NOT NULL COMMENT '交易内容描述',
  `attach` varchar(1024) CHARACTER SET utf8mb4 NOT NULL COMMENT '自定义数据',

  `notify_url` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '回调通知url',
  `notify_result` varchar(1024) CHARACTER SET utf8mb4 NOT NULL COMMENT '回调内容',
  `notify_ts` bigint NOT NULL DEFAULT '0' COMMENT '回调时间',

  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`tsn`),
  UNIQUE KEY `out_trade_no` (`out_trade_no`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_uid` (`uid`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
