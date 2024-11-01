package logic

import (
	"context"

	"fox-server-core/rpc/internal/svc"
	"fox-server-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *GetUserListLogic) GetUserList(in *core.UserListReq) (*core.UserListResp, error) {
	// todo: add your logic here and delete this line

	return &core.UserListResp{}, nil
}
