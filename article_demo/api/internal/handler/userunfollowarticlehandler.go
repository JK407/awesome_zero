package handler

import (
	"net/http"

	"article_demo/api/internal/logic"
	"article_demo/api/internal/svc"
	"article_demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUnFollowArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUnFollowArticleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserUnFollowArticleLogic(r.Context(), svcCtx)
		resp, err := l.UserUnFollowArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
