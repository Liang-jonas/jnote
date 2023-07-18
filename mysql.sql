use JNA;

CREATE TABLE `api`
(
    `id`       bigint       NOT NULL,
    `name`     varchar(100) NOT NULL COMMENT '路由名称',
    `path`     varchar(100) NOT NULL COMMENT '路径',
    `describe` text COMMENT '路由描述',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='后端API表';

CREATE TABLE `front-route`
(
    `id`          bigint       NOT NULL,
    `name`        varchar(100) NOT NULL COMMENT '路由名称',
    `path`        varchar(100) NOT NULL COMMENT '路径',
    `description` text COMMENT '路由描述',
    `index`       int          NOT NULL COMMENT '在同级中显示的位置',
    `type`        tinyint DEFAULT '0' COMMENT '前端路由类型, 真为业务路由,假为工具路由',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='前端路由表';

CREATE TABLE `group`
(
    `id`          bigint NOT NULL,
    `name`        varchar(100) DEFAULT NULL COMMENT '组名称',
    `description` varchar(100) DEFAULT NULL COMMENT '组描述',
    `is-del`      tinyint      DEFAULT '1' COMMENT '是否可以被删除',
    `create-at`   datetime         DEFAULT NULL COMMENT '创建时间',
    `delete-at`   datetime         DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `group_UN` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='组';

CREATE TABLE `image`
(
    `id`          bigint       NOT NULL,
    `name`        varchar(100) NOT NULL COMMENT '图片名称',
    `create-at`   datetime         DEFAULT NULL COMMENT '创建时间',
    `delete-at`   datetime         DEFAULT NULL COMMENT '删除时间',
    `description` varchar(255) DEFAULT NULL COMMENT '图片描述',
    `location`    varchar(255) DEFAULT NULL COMMENT '图片位置',
    `owner`       varchar(36)  DEFAULT NULL COMMENT '图片所有者'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `map-policy-api`
(
    `id`        bigint NOT NULL,
    `policy-id` bigint NOT NULL,
    `api-id`    bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略api映射表';

CREATE TABLE `map-policy-front-route`
(
    `id`              bigint NOT NULL,
    `front-router-id` bigint NOT NULL,
    `policy-id`       bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略前端路由表';

CREATE TABLE `map-policy-image`
(
    `id`        bigint NOT NULL,
    `policy-id` bigint NOT NULL,
    `image-id`  bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略图片映射表';

CREATE TABLE `map-policy-note`
(
    `id`        bigint NOT NULL,
    `note-id`   bigint NOT NULL,
    `policy-id` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略笔记映射表';

CREATE TABLE `note`
(
    `id`        bigint       NOT NULL,
    `title`     varchar(100) DEFAULT NULL COMMENT '笔记标题',
    `owner`     varchar(100) NOT NULL COMMENT '所有者id',
    `create-at` datetime         DEFAULT NULL COMMENT '创建时间',
    `delete-at` datetime         DEFAULT NULL COMMENT '删除时间',
    `update-at` datetime         DEFAULT NULL COMMENT '更新时间',
    UNIQUE KEY `note_UN` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='笔记表';

CREATE TABLE `policy`
(
    `id`       bigint NOT NULL,
    `user-id`  bigint DEFAULT NULL,
    `group-id` bigint DEFAULT NULL,
    UNIQUE KEY `policy_UN` (`id`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='策略组';

CREATE TABLE `theme`
(
    `id`          bigint NOT NULL,
    `name`        varchar(100) DEFAULT NULL COMMENT '主题名称',
    `description` varchar(100) DEFAULT NULL COMMENT '主题描述',
    `context`     varchar(100) DEFAULT NULL COMMENT '主题的内容',
    `create-at`   datetime         DEFAULT NULL COMMENT '创建时间',
    `update-at`   datetime         DEFAULT NULL COMMENT '更新时间',
    `delete-at`   datetime         DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `theme_UN` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='主题表';

CREATE TABLE `user`
(
    `id`        bigint       NOT NULL COMMENT '用户id',
    `username`  varchar(100) NOT NULL COMMENT '用户名',
    `password`  varchar(32)  DEFAULT NULL COMMENT '密码',
    `theme-id`  bigint       DEFAULT NULL COMMENT '主题id',
    `group-id`  bigint       DEFAULT NULL COMMENT '用户组id',
    `disable`   tinyint      DEFAULT '1' COMMENT '是否禁用',
    `avatar`    varchar(100) DEFAULT NULL COMMENT '用户头像',
    `create-at` datetime         DEFAULT NULL COMMENT '创建时间',
    `update-at` datetime         DEFAULT NULL COMMENT '更新时间',
    `delete-at` datetime         DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY `user_UN` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

CREATE TABLE `workspaces`
(
    `id`        bigint       NOT NULL,
    `name`     varchar(100) DEFAULT NULL COMMENT '工作空间名称',
    `owner`     varchar(100) NOT NULL COMMENT '所有者id',
    `create-at` datetime         DEFAULT NULL COMMENT '创建时间',
    `delete-at` datetime         DEFAULT NULL COMMENT '删除时间',
    `update-at` datetime         DEFAULT NULL COMMENT '更新时间',
    UNIQUE KEY `note_UN` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='工作空间表';
