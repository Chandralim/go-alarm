create table `permission_group_detail` (
    `ordinal` int not null,
    `p_change` tinyint(1) not null default '0',
    `permission_group_id` bigint unsigned not null,
    `permission_list_name` varchar(255) not null,
    `created_user` bigint unsigned not null,
    `updated_user` bigint unsigned not null,
    `created_at` timestamp null,
    `updated_at` timestamp null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `permission_group_detail`
add constraint `permission_group_detail_permission_group_id_foreign` foreign key (`permission_group_id`) references `permission_group` (`id`) on delete restrict on update cascade;
alter table `permission_group_detail`
add constraint `permission_group_detail_created_user_foreign` foreign key (`created_user`) references `users` (`id`) on delete restrict on update cascade;
alter table `permission_group_detail`
add constraint `permission_group_detail_updated_user_foreign` foreign key (`updated_user`) references `users` (`id`) on delete restrict on update cascade;