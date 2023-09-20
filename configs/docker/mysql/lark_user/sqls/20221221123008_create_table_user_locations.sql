DROP TABLE IF EXISTS `user_locations`;
CREATE TABLE `user_locations` (
  `uid` bigint unsigned NOT NULL COMMENT '用户ID',
  `longitude` DECIMAL(10,6) NOT NULL DEFAULT '0.000000' COMMENT '经度',
  `latitude` DECIMAL(10,6) NOT NULL DEFAULT '0.000000' COMMENT '纬度',
  `online_ts` bigint NOT NULL DEFAULT '0' COMMENT '最后一次上线时间',
  `created_ts` bigint NOT NULL DEFAULT '0',
  `updated_ts` bigint NOT NULL DEFAULT '0',
  `deleted_ts` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;