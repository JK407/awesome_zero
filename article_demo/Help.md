查询订阅最高的文章内容
用户订阅文章
用户取订阅文章


CREATE TABLE users (
user_id INT AUTO_INCREMENT PRIMARY KEY,
username VARCHAR(50) NOT NULL,
email VARCHAR(100) NOT NULL,
UNIQUE(email)
);


CREATE TABLE articles (
article_id  INT AUTO_INCREMENT PRIMARY KEY,
title VARCHAR(200) NOT NULL,
content TEXT,
author VARCHAR(100),
);

CREATE TABLE user_subscriptions (
subscription_id  INT AUTO_INCREMENT PRIMARY KEY,
user_id INT NOT NULL,
article_id INT NOT NULL,
FOREIGN KEY (user_id) REFERENCES users(user_id),
FOREIGN KEY (article_id) REFERENCES articles(article_id),
UNIQUE(user_id, article_id)
);
