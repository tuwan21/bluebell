CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
            CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `community_id` int(10) unsigned NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
        CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`community_id`),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

insert into `community`(id,community_id,community_name,introduction) VALUES ('1','1','Go','Golang');
insert into `community`(id,community_id,community_name,introduction) VALUES ('2','2','leetcode','LeetCode');
insert into `community`(id,community_id,community_name,introduction) VALUES ('3','3','CSGO','Rush B');
insert into `community`(id,community_id,community_name,introduction) VALUES ('4','4','LOL','!');


DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) unsigned NOT NULL,
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL,
    `author_id` bigint(20) NOT NULL,
    `community_id` bigint(20) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
                CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;