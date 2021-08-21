package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) GradeLogic {
	return GradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeLogic) Grade(req types.VoteReq) (*types.VoteResp, error) {
	// todo: add your logic here and delete this line

	return &types.VoteResp{}, nil
}
