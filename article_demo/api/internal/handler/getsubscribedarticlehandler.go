package handler

import (
	"net/http"

	"article_demo/api/internal/logic"
	"article_demo/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSubscribedArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetSubscribedArticleLogic(r.Context(), svcCtx)
		resp, err := l.GetSubscribedArticle()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
