
use go_zero;
CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'user id',
    `name` varchar(255) NOT NULL COMMENT 'user name',
    `description` varchar(255) NOT NULL COMMENT 'user description',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


# 插入数据
INSERT INTO `user` (`name`, `description`) VALUES ('user1', 'user1 description');
INSERT INTO `user` (`name`, `description`) VALUES ('user2', 'user2 description');
INSERT INTO `user` (`name`, `description`) VALUES ('user3', 'user3 description');
INSERT INTO `user` (`name`, `description`) VALUES ('user4', 'user4 description');
INSERT INTO `user` (`name`, `description`) VALUES ('user5', 'user5 description');