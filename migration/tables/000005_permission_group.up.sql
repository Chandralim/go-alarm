create table `permission_group` (
    `id` bigint unsigned not null auto_increment primary key,
    `name` varchar(30) not null,
    `created_user` bigint unsigned not null,
    `updated_user` bigint unsigned not null,
    `created_at` timestamp null,
    `updated_at` timestamp null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `permission_group`
add constraint `permission_group_created_user_foreign` foreign key (`created_user`) references `is_users` (`id`) on delete restrict on update cascade;
alter table `permission_group`
add constraint `permission_group_updated_user_foreign` foreign key (`updated_user`) references `is_users` (`id`) on delete restrict on update cascade;
alter table `permission_group`
add unique `permission_group_name_unique`(`name`);