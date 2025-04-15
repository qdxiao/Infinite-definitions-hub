package template

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infni-backend/apps/api/internal/logic/template"
	"infni-backend/apps/api/internal/svc"
	"infni-backend/apps/api/internal/types"
)

func HandlerNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := template.NewHandlerNameLogic(r.Context(), svcCtx)
		resp, err := l.HandlerName(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
