package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ActivitiesInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivitiesInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) ActivitiesInfoLogic {
	return ActivitiesInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivitiesInfoLogic) ActivitiesInfo(req types.Actid) (*types.ActivityResp, error) {
	// todo: add your logic here and delete this line

	return &types.ActivityResp{}, nil
}
