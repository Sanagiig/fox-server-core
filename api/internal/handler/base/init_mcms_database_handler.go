package base

import (
	"net/http"

	"fox-server-core/api/internal/logic/base"
	"fox-server-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InitMcmsDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewInitMcmsDatabaseLogic(r.Context(), svcCtx)
		resp, err := l.InitMcmsDatabase()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
