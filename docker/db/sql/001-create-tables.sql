---- drop ----
DROP TABLE IF EXISTS `samples`;

---- create ----
create table IF not exists `samples`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);
---- insert ----
INSERT INTO samples values (1,'sample',CURRENT_TIMESTAMP(),CURRENT_TIMESTAMP(),NULL)

---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users`
(
 `user_id`          INT(20) AUTO_INCREMENT,
 `user_nm`          VARCHAR(20) NOT NULL,
 `password`         VARCHAR(20) NOT NULL,
 `email`            VARCHAR(255) NOT NULL,
    PRIMARY KEY (`user_id`)
);
---- insert ----
INSERT INTO users values (1,'ゲスト','guest_pass','guest@email.com')
