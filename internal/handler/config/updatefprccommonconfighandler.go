package config

import (
	"net/http"

	"easy-frpc/internal/logic/config"
	"easy-frpc/internal/svc"
	"easy-frpc/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateFprcCommonConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommonConfigRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := config.NewUpdateFprcCommonConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateFprcCommonConfig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
