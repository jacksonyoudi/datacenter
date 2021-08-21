package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnrollLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnrollLogic {
	return EnrollLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnrollLogic) Enroll(req types.EnrollReq) (*types.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &types.EnrollResp{}, nil
}
