drop table if exists `message`;
create table if not exists `message` (
    `id`  bigint unsigned NOT NULL auto_increment COMMENT '用户ID 系统生成',
    `created_at` datetime(3) not null,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime(3) DEFAULT NULL,
    `form_id` bigint not null comment '消息发送方',
    `target_id` bigint not null comment '消息接收方',
    `type` int(4) not null comment '消息种类',
    `content` varchar(200) default '' comment '消息内容',
    `content_type` int not null comment '消息类型',
    `avatar` varchar(200) default null comment '头像URL',
    `desc` varchar(200) default null comment '消息描述',
    `extends` varchar(200) default null comment '拓展字段',
     primary key (`id`),
     unique key `uni_message_form_id` (`form_id`, `type`),
     unique key `uni_message_target_id` (`target_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;