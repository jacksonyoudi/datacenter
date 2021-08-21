package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleInitLogic {
	return ArticleInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleInitLogic) ArticleInit(req types.SearchReq) (*types.SearchResp, error) {
	// todo: add your logic here and delete this line

	return &types.SearchResp{}, nil
}
