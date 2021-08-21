package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ActivityInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivityInfoLogic {
	return ActivityInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityInfoLogic) ActivityInfo(req types.Actid) (*types.ActivityResp, error) {
	// todo: add your logic here and delete this line

	return &types.ActivityResp{}, nil
}
