CREATE TABLE `user_tab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `user_name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '加密后的密码',
  `email` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `salt` char(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '盐',
  `profile_uri` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像图片存储地址',
  `user_type` tinyint unsigned NOT NULL COMMENT '用户类型 0:买家 1:卖家',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `update_time` int unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_name_type` (`user_name`,`user_type`)
) ENGINE=InnoDB AUTO_INCREMENT=110007 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表'


CREATE TABLE `seller_shop_tab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺名',
  `shop_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺唯一id',
  `introduction` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺简介',
  `user_name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺所属用户名',
  `level` tinyint NOT NULL COMMENT '店铺等级',
  `location` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺地址',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `update_time` int unsigned NOT NULL COMMENT '修改',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_shop_id` (`shop_id`),
  UNIQUE KEY `index_name` (`name`),
  KEY `idx_user_name` (`user_name`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB AUTO_INCREMENT=642836 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='卖家店铺表'

CREATE TABLE `product_tab` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `product_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品唯一id',
  `title` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品标题',
  `shop_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '店铺所属用户id',
  `cover_uri` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '封面图',
  `price` int unsigned NOT NULL COMMENT '价格',
  `stock` int unsigned NOT NULL COMMENT '库存',
  `brand_id` bigint NOT NULL COMMENT '品牌id',
  `category_id` bigint NOT NULL COMMENT '商品类目id',
  `show_uris` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '展示图',
  `details` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商品详情：暂时图片组; todo详情内容较多',
  `status` int unsigned NOT NULL COMMENT '状态: 1:正常 2:审核中 3:下架 4:禁止 5:编辑 6:删除',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `update_time` int unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
--  UNIQUE KEY `uniq_product_id` (`product_id`),
  UNIQUE KEY `uniq_shop_product` (`shop_id`,`product_id`),
  KEY `idx_shop_product_status` (`shop_id`,`status`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB AUTO_INCREMENT=235137 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品信息表'


CREATE TABLE `product_attr_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `product_id` varchar(36) NOT NULL COMMENT '商品id',
  `name` varchar(20) NOT NULL COMMENT '属性名称',
  `value` varchar(32) NOT NULL COMMENT '属性值',
  `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_product_attr`  (`product_id`,`name`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品属性表';


CREATE TABLE `category_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `name` varchar(20) NOT NULL COMMENT '类目名称',
  `parent_id` bigint NOT NULL COMMENT '类目父id',
  `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`),
  KEY `idx_create_time` (`create_time`),
  KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='类目信息表';

CREATE TABLE `attr_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `name` varchar(20)  NOT NULL COMMENT '属性名称',
  `desc` varchar(128) NOT NULL COMMENT '属性描述',
  `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name`  (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='属性描述表';

CREATE TABLE `brand_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `name` varchar(20) NOT NULL COMMENT '品牌名称',
  `desc` varchar(128) NOT NULL COMMENT '品牌描述',
  `log_uri` varchar(64) NOT NULL COMMENT '品牌log图',
  `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='品牌信息表';
