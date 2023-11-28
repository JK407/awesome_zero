use go_zero;
CREATE TABLE users
(
    user_id  INT  NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    username VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    email    VARCHAR(100) NOT NULL UNIQUE DEFAULT '' COMMENT '邮箱',
    PRIMARY KEY(`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE articles
(
    article_id INT NOT NULL  AUTO_INCREMENT  COMMENT '文章ID',
    title      VARCHAR(200) NOT NULL DEFAULT '' COMMENT '标题',
    content    TEXT  NOT NULL DEFAULT '' COMMENT '内容',
    author     VARCHAR(100) DEFAULT '' COMMENT '作者',
    PRIMARY KEY(`article_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE user_subscriptions
(
    subscription_id INT NOT NULL AUTO_INCREMENT COMMENT '订阅ID',
    user_id         INT NOT NULL DEFAULT 0 COMMENT '用户ID',
    article_id      INT NOT NULL DEFAULT 0 COMMENT '文章ID',
    PRIMARY KEY(`subscription_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



INSERT INTO users (username, email) VALUES
                                        ('user1', 'user1@example.com'),
                                        ('user2', 'user2@example.com');
INSERT INTO articles (title, content, author) VALUES
                                                  ('Article 1', 'Content for Article 1', 'Author 1'),
                                                  ('Article 2', 'Content for Article 2', 'Author 2');
-- 让用户1订阅文章1和文章2
INSERT INTO user_subscriptions (user_id, article_id) VALUES
                                                         (1, 1),
                                                         (1, 2);

-- 让用户2订阅文章1
INSERT INTO user_subscriptions (user_id, article_id) VALUES
    (2, 1);



select a.article_id, a.title, a.content from articles as a join (select article_id,Count(*) as subscription_count from user_subscriptions group by article_id order by subscription_count desc) sub on a.article_id = sub.article_id


select `subscription_id`,`user_id`,`article_id` from `user_subscriptions` where `user_id` = 1 and `article_id` = 1 limit 1