package config

import (
	"context"

	"easy-frpc/internal/svc"
	"easy-frpc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrpcConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFrpcConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrpcConfigLogic {
	return &GetFrpcConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFrpcConfigLogic) GetFrpcConfig() (resp *types.ConfigResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
