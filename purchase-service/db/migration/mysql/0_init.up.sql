START TRANSACTION;

CREATE TABLE IF NOT EXISTS `subscription_plans` (
    `tier` enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM') NOT NULL,
    `currency_symbol` varchar(64) NOT NULL,
    `price` decimal(23,8) NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`tier`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii;

CREATE TABLE IF NOT EXISTS `user_tiers` (
    `user_id` bigint unsigned NOT NULL,
    `tier` enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM') NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`user_id`),
    KEY `idx_user_tiers_tier` (`tier`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`user_id`)
    PARTITIONS 16;

CREATE TABLE IF NOT EXISTS `voucher_pools` (
    `code` varchar(64)  NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii;

CREATE TABLE IF NOT EXISTS `voucher_configurations` (
    `id` bigint unsigned NOT NULL,
    `calculation_type` enum('VCT_PERCENTAGE','VCT_AMOUNT') NOT NULL,
    `currency_symbol` varchar(64) NOT NULL,
    `value` decimal(23,8) NOT NULL,
    `max_qty` INT unsigned NOT NULL,
    `allocated_qty` INT unsigned NOT NULL,
    `redeemed_qty` INT unsigned NOT NULL,
    `start_at` bigint unsigned NOT NULL,
    `end_at` bigint unsigned NOT NULL,
    `status` enum('VS_DRAFT','VS_AVAILABLE','VS_UNAVAILABLE') NOT NULL,
    `applied_tier` enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM') NOT NULL,
    `type` enum('VGT_CAMPAIGN') NOT NULL,
    `campaign_id` bigint unsigned NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_voucher_configurations_calculation_type` (`calculation_type`),
    KEY `idx_voucher_configurations_start_at` (`start_at`),
    KEY `idx_voucher_configurations_end_at` (`end_at`),
    KEY `idx_voucher_configurations_status` (`status`),
    KEY `idx_voucher_configurations_applied_tier` (`applied_tier`),
    KEY `idx_voucher_configurations_campaign_id` (`campaign_id`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`id`)
    PARTITIONS 16;

CREATE TABLE IF NOT EXISTS `user_vouchers` (
    `id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `voucher_configuration_id` bigint unsigned NOT NULL,
    `voucher_code` varchar(64)  NOT NULL,
    `expired_at` bigint unsigned NOT NULL,
    `status` enum('UVS_ALLOCATED','UVS_USED','UVS_EXPIRED') NOT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_user_vouchers_user_id` (`user_id`),
    KEY `idx_user_vouchers_voucher_configuration_id` (`voucher_configuration_id`),
    KEY `idx_user_vouchers_expired_at` (`expired_at`),
    KEY `idx_user_vouchers_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`id`)
    PARTITIONS 16;

CREATE TABLE IF NOT EXISTS `transactions` (
    `id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `currency_symbol` varchar(64) NOT NULL,
    `original_amount` decimal(23,8) NOT NULL,
    `discount_amount` decimal(23,8) NOT NULL,
    `actual_amount` decimal(23,8) NOT NULL,
    `subscription_plan_tier` enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM') DEFAULT NULL,
    `user_voucher_id` bigint unsigned DEFAULT NULL,
    `created_at` bigint unsigned DEFAULT NULL,
    `updated_at` bigint unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_transactions_user_id` (`user_id`),
    KEY `idx_transactions_subscription_plan_tier` (`subscription_plan_tier`),
    KEY `idx_transactions_user_voucher_id` (`user_voucher_id`)
) ENGINE=InnoDB DEFAULT CHARSET=ascii
    PARTITION BY LINEAR HASH (`id`)
    PARTITIONS 16;

COMMIT;
