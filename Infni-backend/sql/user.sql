CREATE DATABASE infinite_community;
USE infinite_community;
CREATE TABLE `user` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` VARCHAR(12) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户密码（加密存储）',
    `account` VARCHAR(15) NOT NULL DEFAULT '' COMMENT '用户账号（登录用）',
    `avator` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像URL',
    `birthday` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '用户生日',
    `profile` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '用户个人简介',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
    `delete_time` DATETIME DEFAULT NULL COMMENT '记录删除时间（软删除用）',
    `is_vip` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否为VIP用户',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_account` (`account`),
    KEY `idx_username` (`username`),
    KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
