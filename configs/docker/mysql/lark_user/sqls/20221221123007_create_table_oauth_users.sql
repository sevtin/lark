DROP TABLE IF EXISTS `oauth_users`;
CREATE TABLE `oauth_users` (
  `oauth_id` bigint unsigned NOT NULL COMMENT '唯一ID',
  `channel` tinyint NOT NULL DEFAULT '0' COMMENT '1:Github 2:Google',
  `openid` varchar(50) NOT NULL DEFAULT '' COMMENT '第三方用户ID',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'lark uid',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '第三方username',
  `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '第三方nickname',
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT 'Email',
  `access_token` varchar(500) NOT NULL COMMENT '第三方AccessToken',
  `refresh_token` varchar(500) NOT NULL COMMENT '第三方RefreshToke',
  `expire` int NOT NULL DEFAULT '0' COMMENT '过期时间 时间戳',
  `avatar_url` varchar(128) NOT NULL COMMENT '第三方头像url',
  `scope` varchar(128) NOT NULL COMMENT '用户授权的作用域，使用逗号（,）分隔',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`oauth_id`),
  KEY `idx_channel_openid_uid` (`channel`,`openid`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;