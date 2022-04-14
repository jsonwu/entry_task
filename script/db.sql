




CREATE TABLE `product_tab` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
  `product_id` varchar(36) NOT NULL COMMENT '商品唯一id',
  `title` varchar(20) NOT NULL COMMENT '商品标题',
  `shop_id` varchar(36) NOT NULL COMMENT '店铺所属用户id',
  `cover_uri` varchar(64)  NOT NULL COMMENT '封面图',
  `price` int(10) unsigned NOT NULL COMMENT '价格',
  `stock` int(10) unsigned NOT NULL COMMENT '库存',
  `brand_id` bigint(20)  NOT NULL COMMENT '品牌id',
  `category_id` bigint(20) NOT NULL COMMENT '商品类目id',
   `show_uris` varchar(512) NOT NULL COMMENT '展示图',
  `details` varchar(1024) NOT NULL COMMENT '商品详情：暂时图片组; todo详情内容较多',
  `status` int unsigned NOT NULL COMMENT '状态: 1:正常 2:审核中 3:下架 4:禁止 5:编辑 6:删除',
  `create_time` int(10) unsigned NOT NULL COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_product_id` (`product_id`),
  UNIQUE KEY `uniq_shop_product` (`shop_id`,`product_id`),
  KEY `idx_shop_product_status` (`shop_id`,`status`),
  KEY `idx_create_time` (`create_time`),
    KEY `idx_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商品信息表';


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
