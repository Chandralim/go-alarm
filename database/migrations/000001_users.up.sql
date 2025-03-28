create table `users` (
  `id` bigint unsigned not null auto_increment primary key,
  `username` varchar(50) not null,
  `password` varchar(255) not null,
  `division` varchar(50) not null,
  `role` varchar(50) not null,
  `is_active` tinyint(1) not null default '1',
  `created_user` bigint unsigned null,
  `updated_user` bigint unsigned null,
  `created_at` timestamp null,
  `updated_at` timestamp null
) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `users`
add constraint `users_created_user_foreign` foreign key (`created_user`) references `users` (`id`) on delete cascade on update cascade;
alter table `users`
add constraint `users_updated_user_foreign` foreign key (`updated_user`) references `users` (`id`) on delete cascade on update cascade;
alter table `users`
add unique `users_username_unique`(`username`);