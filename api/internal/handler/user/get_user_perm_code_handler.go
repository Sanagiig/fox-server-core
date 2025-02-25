package user

import (
	"net/http"

	"fox-server-core/api/internal/logic/user"
	"fox-server-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserPermCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserPermCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPermCode()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
