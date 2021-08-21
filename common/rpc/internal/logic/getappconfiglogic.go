package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/common/rpc/common"
	"github.com/jacksonyoudi/datacenter/common/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAppConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppConfigLogic {
	return &GetAppConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppConfigLogic) GetAppConfig(in *common.AppConfigReq) (*common.AppConfigResp, error) {
	// todo: add your logic here and delete this line

	return &common.AppConfigResp{}, nil
}
