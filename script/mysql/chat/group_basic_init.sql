drop table if exists `group_basic`;
create table if not exists `group_basic` (
    `id`  bigint unsigned NOT NULL COMMENT '用户ID 系统生成',
    `created_at` datetime(3) not null,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime(3) DEFAULT NULL,
    `name` varchar(20) NOT NULL COMMENT 'group名称',
    `owner_id` bigint(20) NOT NULL COMMENT '群主userId',
    `icon` varchar(200) DEFAULT '' COMMENT 'group icon',
    `desc` varchar(200) DEFAULT '' COMMENT '消息内容',
    `extends` varchar(200) DEFAULT '' COMMENT '拓展字段',
    primary key (`id`),
    UNIQUE KEY `uni_group_basic_owner_id` (`owner_id`),
    UNIQUE KEY `uni_group_basic_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;