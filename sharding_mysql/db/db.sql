create database apporder default character set utf8mb4 collate utf8mb4_unicode_ci;
use apporder;

create database appuser default character set utf8mb4 collate utf8mb4_unicode_ci;
use appuser;
CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255)  NOT NULL DEFAULT '',
    `password` varchar(255) NOT NULL DEFAULT '',
    `email` varchar(255) NOT NULL DEFAULT '',
    `phone` varchar(255) NOT NULL DEFAULT '',
    `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'status:0 default, 1 active, 2 disable',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_username_unique` (`username`),
    KEY `users_email_index` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
create table `users_0001` select * from users where 1=2;
