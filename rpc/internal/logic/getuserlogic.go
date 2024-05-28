package logic

import (
	"context"

	"im/rpc/internal/svc"
	"im/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	if u, ok := users[in.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return &user.GetUserResp{}, nil
}

// data.go

type User struct {
	Id    string
	Name  string
	Phone string
	Pass  string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "高瑞克",
		Phone: "13100001111",
		Pass:  "123456",
	},
	"2": {
		Id:    "2",
		Name:  "rick",
		Phone: "13200001111",
		Pass:  "123456",
	},
}
