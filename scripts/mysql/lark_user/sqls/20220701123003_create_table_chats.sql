DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats` (
  `chat_id` bigint unsigned NOT NULL COMMENT 'chat ID',
  `creator_uid` bigint unsigned NOT NULL COMMENT '创建者 uid',
  `chat_type` tinyint(2) unsigned NOT NULL COMMENT 'chat type 1:私聊/2:群聊',
  `avatar` varchar(160) NOT NULL COMMENT '小图 72*72',
  `name` varchar(128) DEFAULT '' COMMENT 'chat标题',
  `about` varchar(255) DEFAULT '' COMMENT '关于',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`chat_id`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_creatorUid` (`creator_uid`),
  KEY `idx_chatType` (`chat_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;