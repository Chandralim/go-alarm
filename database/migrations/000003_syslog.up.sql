create table `syslog` (
    `created_at` timestamp default CURRENT_TIMESTAMP not null,
    `ip_address` varchar(25) not null,
    `created_user` bigint unsigned null,
    `module` varchar(50) not null,
    `module_id` bigint null,
    `action` varchar(20) not null,
    `note` longtext not null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `syslog`
add constraint `syslog_created_user_foreign` foreign key (`created_user`) references `users` (`id`) on delete restrict on update cascade;