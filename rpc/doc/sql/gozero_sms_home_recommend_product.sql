create table sms_home_recommend_product
(
    id               bigint auto_increment
        primary key,
    product_id       bigint      null,
    product_name     varchar(64) null,
    recommend_status int(1)      null,
    sort             int(1)      null
)
    comment '人气推荐商品表';

INSERT INTO gozero.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort) VALUES (3, 26, '华为 HUAWEI P20 ', 1, 0);
INSERT INTO gozero.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort) VALUES (4, 27, '小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待', 1, 0);
INSERT INTO gozero.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort) VALUES (5, 28, '小米 红米5A 全网通版 3GB+32GB 香槟金 移动联通电信4G手机 双卡双待', 1, 0);
INSERT INTO gozero.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort) VALUES (6, 29, 'Apple iPhone 8 Plus 64GB 红色特别版 移动联通电信4G手机', 1, 0);
INSERT INTO gozero.sms_home_recommend_product (id, product_id, product_name, recommend_status, sort) VALUES (7, 30, 'HLA海澜之家简约动物印花短袖T恤', 1, 100);