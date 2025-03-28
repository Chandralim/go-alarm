create table `permission_list` (`name` varchar(255) not null) default character set utf8mb4 collate 'utf8mb4_unicode_ci';
alter table `permission_list`
add unique `permission_list_name_unique`(`name`);