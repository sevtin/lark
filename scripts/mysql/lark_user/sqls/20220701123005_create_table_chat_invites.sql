DROP TABLE IF EXISTS `chat_invites`;
CREATE TABLE `chat_invites` (
  `invite_id` bigint unsigned NOT NULL COMMENT 'invite ID',
  `chat_id` bigint unsigned NOT NULL COMMENT 'Chat ID',
  `chat_type` tinyint(2) unsigned NOT NULL COMMENT '1:私聊/2:群聊',
  `initiator_uid` bigint unsigned NOT NULL COMMENT '发起人 UID',
  `invitee_uid` bigint unsigned NOT NULL COMMENT '被邀请人 UID',
  `invitation_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '邀请消息',
  `handler_uid` bigint unsigned NOT NULL DEFAULT '0'  COMMENT '处理人 UID',
  `handle_result` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '结果',
  `handle_msg` varchar(255) DEFAULT '' COMMENT '处理消息',
  `handled_ts` bigint NOT NULL DEFAULT '0' COMMENT '处理时间',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`invite_id`),
  KEY `id_initiatorUid_handleResult_deletedTs` (`initiator_uid`,`handle_result`,`deleted_ts`),
  KEY `id_inviteeUid_handleResult_deletedTs` (`invitee_uid`,`handle_result`,`deleted_ts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;