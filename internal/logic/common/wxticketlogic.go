package logic

import (
	"context"

	"github.com/jacksonyoudi/datacenter/internal/svc"
	"github.com/jacksonyoudi/datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WxTicketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxTicketLogic {
	return WxTicketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxTicketLogic) WxTicket(req types.SnsReq) (*types.WxShareResp, error) {
	// todo: add your logic here and delete this line

	return &types.WxShareResp{}, nil
}
