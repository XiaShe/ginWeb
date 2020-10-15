CREATE TABLE `user`
(
    `id`          BIGINT(20)                             NOT NULL AUTO_INCREMENT,
    `user_id`     BIGINT(20)                             NOT NULL,
    `username`    VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       VARCHAR(64) COLLATE utf8mb4_general_ci,
    `gender`      TINYINT(4)                             NOT NULL DEFAULT '' 0 '',
    `create_time` DATETIME                               NOT NULL,
    `update_time` TIMESTAMP                              NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`
(
    `id`             INT(11)                                 NOT NULL AUTO_INCREMENT,
    `community_id`   INT(10) UNSIGNED                        NOT NULL,
    `community_name` VARCHAR(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   VARCHAR(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    DATETIME                                NOT NULL,
    `update_time`    TIMESTAMP                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

INSERT INTO `community`
VALUES (''1'', ''1'', ''Go'', ''萌新求关照'', ''2016-11-01 08:10:10'', ''2016-11-01 08:10:10'');
INSERT INTO `community`
VALUES (''2'', ''2'', ''Java'', ''史上最强'', ''2020-01-01 08:00:00'', ''2020-01-01 08:00:00'');
INSERT INTO `community`
VALUES (''3'', ''3'', ''Python'', ''这是我的时代'', ''2018-08-07 08:30:00'', ''2018-08-07 08:30:00'');
INSERT INTO `community`
VALUES (''4'', ''4'', ''C++'', ''你们还是太年轻了'', ''2016-01-01 08:00:00'', ''2016-01-01 08:00:00'');

CREATE TABLE `post`
(
    `id`           BIGINT(20)                               NOT NULL AUTO_INCREMENT,
    `post_id`      BIGINT(20)                               NOT NULL COMMENT '' 帖子id '',
    `title`        VARCHAR(128) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '' 标题 '',
    `content`      VARCHAR(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '' 内容 '',
    `author_id`    BIGINT(20)                               NOT NULL COMMENT '' 作者的用户id '',
    `community_id` BIGINT(20)                               NOT NULL COMMENT '' 所属社区 '',
    `status`       TINYINT(4)                               NOT NULL DEFAULT '' 1 '' COMMENT '' 帖子状态 '',
    `create_time`  DATETIME                                 NOT NULL COMMENT '' 创建时间 '',
    `update_time`  TIMESTAMP                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '' 更新时间 '',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

