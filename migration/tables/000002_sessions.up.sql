create table `sessions` (
    `token` varchar(255) not null,
    `user_id` bigint unsigned not null,
    `created_at` timestamp default CURRENT_TIMESTAMP not null,
    `updated_at` timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `sessions`
add constraint `sessions_user_id_foreign` foreign key (`user_id`) references `is_users` (`id`) on delete restrict on update cascade;
alter table `sessions`
add unique `sessions_token_unique`(`token`);