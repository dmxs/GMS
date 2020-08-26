DROP TABLE IF EXISTS user;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密码',
  `image` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '图片地址',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `role_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_username_uindex` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';
