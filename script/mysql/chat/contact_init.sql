drop table if exists `contact`;
create table if not exists `contact` (
    `id`  bigint unsigned NOT NULL auto_increment COMMENT '用户ID 系统生成',
    `created_at` datetime(3) not null,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime(3) DEFAULT NULL,
    `name` varchar(20) not null comment 'group名称',
    `owner_id` bigint not null comment 'owner_id',
    `target_id` bigint not null comment 'target_id',
    `type` int(4) not null comment '消息种类',
    `desc` varchar(200) default '' comment '消息内容',
    primary key (`id`),
    unique key `uni_contact_owner_id` (`owner_id`),
    unique key `uni_contact_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;