create TABLE `user_tracking`.`user_visited`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `ip`         VARCHAR(64),
    `browser`    VARCHAR(255),
    `location`   VARCHAR(255),
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    KEY          `idx_updated_at` (`updated_at`),
    KEY          `idx_ip` (`ip`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'all user visited';
