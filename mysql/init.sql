CREATE USER 'user'@'%' IDENTIFIED BY 'password';
CREATE DATABASE go_ecom_user;
GRANT ALL PRIVILEGES ON go_ecom_user.* TO 'user'@'%';
CREATE DATABASE go_ecom_order;
GRANT ALL PRIVILEGES ON go_ecom_order.* TO 'user'@'%';
CREATE DATABASE go_ecom_product;
GRANT ALL PRIVILEGES ON go_ecom_product.* TO 'user'@'%';
create table if not exists go_ecom_order.`order`
(
    id               int auto_increment
        primary key,
    user_id          int            null,
    billing_address  varchar(255)   not null,
    shipping_address varchar(255)   not null,
    total_amount     decimal(10, 2) not null,
    status           varchar(20)    not null,
    created          datetime(6)    not null,
    updated          datetime(6)    not null
);

create table if not exists go_ecom_order.order_product
(
    id         int auto_increment
        primary key,
    order_id   int            null,
    product_id int            null,
    quantity   int            not null,
    price      decimal(10, 2) not null,
    created    datetime(6)    not null,
    updated    datetime(6)    not null,
    constraint order_product_ibfk_1
        foreign key (order_id) references go_ecom_order.`order` (id)
);

create index order_id
    on go_ecom_order.order_product (order_id);

create table if not exists go_ecom_product.product
(
    id          int auto_increment
        primary key,
    name        varchar(255)                         not null,
    description text                                 null,
    price       decimal(10, 2)                       not null,
    image_url   varchar(255)                         null,
    category    varchar(50)                          null,
    is_active   tinyint(1) default 0                 not null,
    created     datetime   default CURRENT_TIMESTAMP not null,
    updated     datetime   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

create table if not exists go_ecom_product.product_user_permission
(
    product_id int                                not null,
    user_id    int                                not null,
    permission enum ('View', 'Update', 'Delete')  null,
    created    datetime default CURRENT_TIMESTAMP not null,
    updated    datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    primary key (product_id, user_id),
    constraint product_user_permission_ibfk_1
        foreign key (product_id) references go_ecom_product.product (id)
);

create table if not exists go_ecom_product.product_version
(
    id         int auto_increment
        primary key,
    product_id int                                not null,
    details    json                               null,
    is_active  tinyint(1)                         null,
    created    datetime default CURRENT_TIMESTAMP not null,
    updated    datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

create table if not exists go_ecom_user.user
(
    id            int auto_increment
        primary key,
    username      varchar(50)                        not null,
    password_hash varchar(255)                       not null,
    role          enum ('Admin', 'Guest')            not null,
    created       datetime default CURRENT_TIMESTAMP not null,
    updated       datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint username
        unique (username)
);
INSERT INTO go_ecom_user.user (username, password_hash, role) VALUES ('admin', '$2a$04$LHtfHuyDbvRvOaVnfNhBQew9HekYjQDp7tAEeyF4b8ErKg/XyBzpC', 'Admin');