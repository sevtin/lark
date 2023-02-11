DROP TABLE IF EXISTS `oauth_clients`;
CREATE TABLE `oauth_clients` (
  `client_id` varchar(50) NOT NULL DEFAULT '' COMMENT '客户端id',
  `client_secret` varchar(64) NOT NULL DEFAULT '' COMMENT '应用密钥',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:正常 1:冻结',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`client_id`),
  UNIQUE KEY `client_id` (`client_id`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;