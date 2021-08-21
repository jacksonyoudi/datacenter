package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeLogic {
	return ChangeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeLogic) Change(req types.AnswerReq) (*types.VoteResp, error) {
	// todo: add your logic here and delete this line

	return &types.VoteResp{}, nil
}
