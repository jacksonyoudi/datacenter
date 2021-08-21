package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AwardInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAwardInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AwardInfoLogic {
	return AwardInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AwardInfoLogic) AwardInfo(req types.Actid) (*types.ActivityResp, error) {
	// todo: add your logic here and delete this line

	return &types.ActivityResp{}, nil
}
