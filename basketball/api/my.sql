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


CREATE TABLE `teams` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "队伍ID",
  `name` varchar(191) DEFAULT NULL COMMENT "队伍名称",
  `description` text COMMENT "队伍描述",
  `league_id` bigint DEFAULT NULL COMMENT "联赛ID",
  `status` bigint DEFAULT '1' COMMENT "状态 1:正常 2:禁用",
  `created_at` datetime(3) DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime(3) DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB Comment="队伍表";


CREATE TABLE `players` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "球员ID",
  `name` varchar(191) DEFAULT NULL COMMENT "球员名称",
  `description` text COMMENT "球员描述",
  `team_id` SMALLINT UNSIGNED DEFAULT NULL COMMENT "队伍ID",
  `number` TINYINT UNSIGNED DEFAULT NULL COMMENT "球员号码",
  `position` varchar(191) DEFAULT NULL COMMENT "球员位置",
  `height` SMALLINT UNSIGNED DEFAULT NULL COMMENT "球员身高",
  `weight` TINYINT UNSIGNED DEFAULT NULL COMMENT "球员体重",
  `status` bigint DEFAULT '1' COMMENT "状态 1:正常 2:禁用",
  `created_at` datetime(3) DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime(3) DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB Comment="球员表";

CREATE TABLE `games` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "比赛ID",
  `name` varchar(191) DEFAULT NULL COMMENT "比赛名称",
  `description` text COMMENT "比赛描述",
  `datetime` datetime DEFAULT NULL COMMENT "比赛时间",
  `league_id` SMALLINT UNSIGNED DEFAULT NULL COMMENT "联赛ID",
  `home_team_id` SMALLINT UNSIGNED DEFAULT NULL COMMENT "主队ID",
  `away_team_id` SMALLINT UNSIGNED DEFAULT NULL COMMENT "客队ID",
  `status` bigint DEFAULT '1' COMMENT "状态 1:正常 2:禁用",
  `created_at` datetime(3) DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime(3) DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB Comment="比赛表";

CREATE TABLE `game_stats` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "比赛事件ID",
  `game_id` bigint DEFAULT NULL COMMENT "比赛ID",
  `player_id` bigint DEFAULT NULL COMMENT "球员ID",
  `type` TINYINT(1) UNSIGNED DEFAULT NULL COMMENT "事件类型 1:进球 2:助攻 3:抢断 4:盖帽 5:失球 6:犯规 7:其他",
  `time` time DEFAULT NULL COMMENT "事件时间",
  `description` text COMMENT "事件描述",
  `created_at` datetime(3) DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime(3) DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB Comment="游戏球员表";