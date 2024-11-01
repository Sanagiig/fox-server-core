package public_user

import (
	"net/http"

	"fox-server-core/api/internal/logic/public_user"
	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterBySmsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public_user.NewRegisterBySmsLogic(r.Context(), svcCtx)
		resp, err := l.RegisterBySms(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
