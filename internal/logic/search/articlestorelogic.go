package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleStoreLogic {
	return ArticleStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleStoreLogic) ArticleStore(req types.ArticleReq) (*types.ArticleReq, error) {
	// todo: add your logic here and delete this line

	return &types.ArticleReq{}, nil
}
