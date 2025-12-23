CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "用户ID",
  `username` varchar(191) DEFAULT NULL COMMENT "用户名",
  `password` longtext COMMENT "密码",
  `status` bigint DEFAULT '1',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_username` (`username`)
) ENGINE=InnoDB Comment="管理员表";


 CREATE TABLE `leagues` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "联赛ID",
  `name` varchar(191) DEFAULT NULL COMMENT "联赛名称",
  `description` text COMMENT "联赛描述",
  `status` bigint DEFAULT '1' COMMENT "状态 1:正常 2:禁用",
  `created_at` datetime(3) DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime(3) DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB Comment="联赛表";