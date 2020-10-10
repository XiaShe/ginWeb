CREATE TABLE `user` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT ,
    `user_id` BIGINT(20) NOT NULL ,
    `username` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `email` VARCHAR(64) COLLATE utf8mb4_general_ci ,
    `gender` TINYINT(4) NOT NULL DEFAULT '0' ,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY (`id`) ,
    UNIQUE KEY `idx_username` (`username`) USING BTREE ,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ;



DROP TABLE IF EXISTS `community`;
create table `community` (
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction` varchar(256) collate utf8mb4_general_ci not null,
    `create_time` timestamp not null default current_timestamp,
    `update_time` datetime not null default current_timestamp on update current_timestamp,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;