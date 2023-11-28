package handler

import (
	"net/http"

	"article_demo/api/internal/logic"
	"article_demo/api/internal/svc"
	"article_demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFollowArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFollowArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFollowArticleLogic(r.Context(), svcCtx)
		resp, err := l.UserFollowArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
