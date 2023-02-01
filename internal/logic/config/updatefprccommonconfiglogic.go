package config

import (
	"context"

	"easy-frpc/internal/svc"
	"easy-frpc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFprcCommonConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFprcCommonConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFprcCommonConfigLogic {
	return &UpdateFprcCommonConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFprcCommonConfigLogic) UpdateFprcCommonConfig(req *types.UpdateCommonConfigRequest) (resp *types.ConfigResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
