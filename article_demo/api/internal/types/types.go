// Code generated by goctl. DO NOT EDIT.
package types

type UserFollowArticleRequest struct {
	UserId    int `json:"user_id"`
	ArticleId int `json:"article_id"`
}

type UserUnFollowArticleRequest struct {
	UserId    int `json:"user_id"`
	ArticleId int `json:"article_id"`
}

type GetSubscribedArticle struct {
	ArticleId      int    `json:"article_id"`
	ArticleTitle   string `json:"article_title"`
	ArticleContent string `json:"article_content"`
	ArticleAuthor  string `json:"article_author"`
}

type UserFollowArticle struct {
	SubscriptionId int `json:"subscription_id"`
}

type UserUnFollowArticle struct {
	Result bool `json:"result"`
}

type GetSubscribedArticleResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data []GetSubscribedArticle `json:"data"`
}

type UserFollowArticleResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data UserFollowArticle `json:"data"`
}

type UserUnFollowArticleResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data UserUnFollowArticle `json:"data"`
}