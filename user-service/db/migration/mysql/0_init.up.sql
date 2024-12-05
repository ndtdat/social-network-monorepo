START TRANSACTION;

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint unsigned NOT NULL,
    `email`  varchar(320) NOT NULL,
    `password` varchar(256) NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`id`)
    PARTITIONS 16;

CREATE TABLE IF NOT EXISTS `campaigns` (
    `id` bigint unsigned NOT NULL,
    `code`  varchar(64) NOT NULL,
    `status` enum('CS_AVAILABLE','CS_UNAVAILABLE') NOT NULL,
    `start_at` bigint unsigned NOT NULL,
    `end_at` bigint unsigned NOT NULL,
    `max_qty` INT unsigned NOT NULL,
    `joined_qty` INT unsigned NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_campaigns_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`id`)
    PARTITIONS 16;

CREATE TABLE IF NOT EXISTS `user_campaigns` (
    `user_id` bigint unsigned NOT NULL,
    `campaign_id`  bigint unsigned NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`user_id`,`campaign_id`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`user_id`)
    PARTITIONS 16;

COMMIT;
