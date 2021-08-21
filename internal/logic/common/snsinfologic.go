package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SnsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSnsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) SnsInfoLogic {
	return SnsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SnsInfoLogic) SnsInfo(req types.SnsReq) (*types.SnsResp, error) {
	// todo: add your logic here and delete this line

	return &types.SnsResp{}, nil
}
