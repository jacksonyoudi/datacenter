package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EnrollListsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnrollListsLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnrollListsLogic {
	return EnrollListsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnrollListsLogic) EnrollLists(req types.Actid) (*types.EnrollResp, error) {
	// todo: add your logic here and delete this line

	return &types.EnrollResp{}, nil
}
