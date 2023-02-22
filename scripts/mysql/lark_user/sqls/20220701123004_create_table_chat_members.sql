DROP TABLE IF EXISTS `chat_members`;
CREATE TABLE `chat_members` (
  `chat_id` bigint unsigned NOT NULL COMMENT 'chat ID',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `chat_type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'chat type 1:私聊/2:群聊',
  `chat_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '名称',
  `remark` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '备注',
  `owner_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '归属人ID',
  `role_id` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `alias` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '别名',
  `member_avatar_key` varchar(50) NOT NULL DEFAULT '' COMMENT 'member头像 72*72',
  `chat_avatar_key` varchar(50) NOT NULL DEFAULT '' COMMENT 'chat头像 72*72',
  `sync` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否同步用户信息 0:同步 1:不同步',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'NORMAL:正常模式 MUTE:开启免打扰 BANNED:被禁言',
  `join_source` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '加入源',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`chat_id`,`uid`),
  UNIQUE KEY `chatId_uid_deletedTs` (`chat_id`,`uid`,`deleted_ts`),
  KEY `idx_deletedTs` (`deleted_ts`),
  KEY `idx_chatType` (`chat_type`),
  KEY `idx_uid` (`uid`),
  KEY `idx_sync` (`sync`),
  KEY `idx_status` (`status`),
  KEY `idx_joinSource` (`join_source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*
 修改`users`.`nickname`,当 `chat_members`.`sync`=0 ,需要同步修改`chat_members`.`alias`
 修改`user_avatars`.`avatar_*`,当 `chat_members`.`sync`=0 ,需要同步修改`chat_members`.`avatar_key`
 需要更新缓存信息
 */
