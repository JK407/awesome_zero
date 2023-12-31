// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articlesFieldNames          = builder.RawFieldNames(&Articles{})
	articlesRows                = strings.Join(articlesFieldNames, ",")
	articlesRowsExpectAutoSet   = strings.Join(stringx.Remove(articlesFieldNames, "`article_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	articlesRowsWithPlaceHolder = strings.Join(stringx.Remove(articlesFieldNames, "`article_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheArticlesArticleIdPrefix = "cache:articles:articleId:"
)

type (
	articlesModel interface {
		Insert(ctx context.Context, data *Articles) (sql.Result, error)
		FindOne(ctx context.Context, articleId int64) (*Articles, error)
		Update(ctx context.Context, data *Articles) error
		Delete(ctx context.Context, articleId int64) error
		GetSubscribedArticles(ctx context.Context) ([]*Articles, error)	 //  获取订阅文章
	}

	defaultArticlesModel struct {
		sqlc.CachedConn
		table string
	}

	Articles struct {
		ArticleId int64          `db:"article_id"` // 文章ID
		Title     string         `db:"title"`      // 标题
		Content   string `db:"content"`    // 内容
		Author    string         `db:"author"`     // 作者
	}
)

func newArticlesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultArticlesModel {
	return &defaultArticlesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`articles`",
	}
}

func (m *defaultArticlesModel) Delete(ctx context.Context, articleId int64) error {
	articlesArticleIdKey := fmt.Sprintf("%s%v", cacheArticlesArticleIdPrefix, articleId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `article_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, articleId)
	}, articlesArticleIdKey)
	return err
}

func (m *defaultArticlesModel) FindOne(ctx context.Context, articleId int64) (*Articles, error) {
	articlesArticleIdKey := fmt.Sprintf("%s%v", cacheArticlesArticleIdPrefix, articleId)
	var resp Articles
	err := m.QueryRowCtx(ctx, &resp, articlesArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", articlesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, articleId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultArticlesModel) Insert(ctx context.Context, data *Articles) (sql.Result, error) {
	articlesArticleIdKey := fmt.Sprintf("%s%v", cacheArticlesArticleIdPrefix, data.ArticleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, articlesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Content, data.Author)
	}, articlesArticleIdKey)
	return ret, err
}

func (m *defaultArticlesModel) Update(ctx context.Context, data *Articles) error {
	articlesArticleIdKey := fmt.Sprintf("%s%v", cacheArticlesArticleIdPrefix, data.ArticleId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `article_id` = ?", m.table, articlesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Content, data.Author, data.ArticleId)
	}, articlesArticleIdKey)
	return err
}

func (m *defaultArticlesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheArticlesArticleIdPrefix, primary)
}

func (m *defaultArticlesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", articlesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticlesModel) tableName() string {
	return m.table
}

func (m *defaultArticlesModel) GetSubscribedArticles(ctx context.Context) ([]*Articles, error) {
	query := fmt.Sprintf("SELECT a.article_id, a.title, a.content, a.author FROM %s AS a JOIN (SELECT article_id, COUNT(*) AS subscription_count FROM user_subscriptions GROUP BY article_id ORDER BY subscription_count DESC) sub ON a.article_id = sub.article_id", m.table)
	var resp []*Articles
		err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


