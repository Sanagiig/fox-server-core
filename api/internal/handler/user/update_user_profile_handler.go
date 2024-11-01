package user

import (
	"net/http"

	"fox-server-core/api/internal/logic/user"
	"fox-server-core/api/internal/svc"
	"fox-server-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProfileInfo
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
