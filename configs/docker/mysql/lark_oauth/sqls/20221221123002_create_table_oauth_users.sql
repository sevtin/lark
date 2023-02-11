CREATE TABLE `oauth_users` (
  `oauth_uid` varchar(50) NOT NULL COMMENT '绑定账号的id',
  `username` varchar(255) NOT NULL COMMENT 'username',
  `access_token` varchar(255) NOT NULL COMMENT 'token',
  `email` varchar(64) NOT NULL COMMENT 'email',
  `company` varchar(64) NOT NULL COMMENT '公司',
  `avatar_url` varchar(128) NOT NULL COMMENT '头像',
  `home_url` varchar(255) NOT NULL COMMENT '主页',
  `created_at` datetime NOT NULL DEFAULT '1970-01-01 08:00:00' COMMENT '账号创建时间',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;