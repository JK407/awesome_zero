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
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsersUserIdPrefix = "cache:users:userId:"
	cacheUsersEmailPrefix  = "cache:users:email:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*Users, error)
		FindOneByEmail(ctx context.Context, email string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		UserId   int64  `db:"user_id"`  // 用户ID
		Username string `db:"username"` // 用户名
		Email    string `db:"email"`    // 邮箱
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, userId int64) error {
	data, err := m.FindOne(ctx, userId)
	if err != nil {
		return err
	}

	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, userId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, userId)
	}, usersEmailKey, usersUserIdKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, userId int64) (*Users, error) {
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, userId)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", usersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

func (m *defaultUsersModel) FindOneByEmail(ctx context.Context, email string) (*Users, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, email)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.UserId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Email)
	}, usersEmailKey, usersUserIdKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.UserId)
	if err != nil {
		return err
	}

	usersEmailKey := fmt.Sprintf("%s%v", cacheUsersEmailPrefix, data.Email)
	usersUserIdKey := fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Username, newData.Email, newData.UserId)
	}, usersEmailKey, usersUserIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsersUserIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
