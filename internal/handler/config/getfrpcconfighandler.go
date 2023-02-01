package config

import (
	"net/http"

	"easy-frpc/internal/logic/config"
	"easy-frpc/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetFrpcConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := config.NewGetFrpcConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetFrpcConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
