create table `permission_group_user` (
    `ordinal` int not null,
    `p_change` tinyint(1) not null default '0',
    `permission_group_id` bigint unsigned not null,
    `user_id` bigint unsigned not null,
    `created_user` bigint unsigned not null,
    `updated_user` bigint unsigned not null,
    `created_at` timestamp null,
    `updated_at` timestamp null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `permission_group_user`
add constraint `permission_group_user_permission_group_id_foreign` foreign key (`permission_group_id`) references `permission_group` (`id`) on delete restrict on update cascade;
alter table `permission_group_user`
add constraint `permission_group_user_user_id_foreign` foreign key (`user_id`) references `is_users` (`id`) on delete restrict on update cascade;
alter table `permission_group_user`
add constraint `permission_group_user_created_user_foreign` foreign key (`created_user`) references `is_users` (`id`) on delete restrict on update cascade;
alter table `permission_group_user`
add constraint `permission_group_user_updated_user_foreign` foreign key (`updated_user`) references `is_users` (`id`) on delete restrict on update cascade;