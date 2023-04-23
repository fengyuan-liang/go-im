drop table if exists `user_basic`;
create table if not exists `user_basic` (
    `id`  bigint unsigned NOT NULL auto_increment COMMENT '用户ID 系统生成',
    `user_id` bigint(15) unsigned NOT NULL COMMENT '用户id，11位数字',
    `user_number` varchar(20) not null comment '用户外显id',
    `name` varchar(128) not null comment '用户名',
    `age` tinyint unsigned not null comment '年龄',
    `password` varchar(256) not null comment '密码 md5散列值',
    `phone_number` varchar(20) default null comment '电话号码',
    `email` varchar(128) default null comment 'email',
    `identity` varchar(128)  default '-1' comment '身份信息',
    `client_ip` varchar(50) default '' comment '客户端ip地址',
    `client_port` varchar(20) default '' comment '端口',
    `login_time` datetime(3) default null comment '最后登陆时间',
    `heart_beat_time` datetime(3) default null comment '最后心跳时间',
    `logout_time` datetime(3) default null comment '最后登陆时间',
    `is_login` boolean default '0' comment '是否处于登陆状态',
    `device_info` varchar(256) default '' comment '设备信息',
    `salt` varchar(128) not null comment 'md5盐',
    `created_at` datetime(3) not null,
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime(3) DEFAULT NULL,
     primary key (`id`),
     unique key `uniq_user_basic_user_id` (`user_id`),
     key `key_user_basic_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;