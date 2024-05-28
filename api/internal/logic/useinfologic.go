package logic

import (
	"context"
	"im/rpc/userclient"

	"im/api/internal/svc"
	"im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseInfoLogic {
	return &UseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseInfoLogic) UseInfo(req *types.UserReq) (resp *types.UserResp, err error) {

	getUserResp, err := l.svcCtx.User.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
}
