DROP TABLE IF EXISTS `avatars`;
CREATE TABLE `avatars` (
  `owner_id` bigint unsigned NOT NULL COMMENT '用户ID/ChatID',
  `owner_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '1:用户头像 2:群头像',
  `avatar_small` varchar(64) NOT NULL COMMENT '小图 72*72',
  `avatar_medium` varchar(64) NOT NULL COMMENT '中图 240*240',
  `avatar_large` varchar(64) NOT NULL COMMENT '大图 480*480',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`owner_id`),
  KEY `idx_ownerType` (`owner_type`),
  KEY `idx_deletedTs` (`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
